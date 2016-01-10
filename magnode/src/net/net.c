/**
 * Author :cz cz.theng@gmail.com
 * Licence MIT
 */

#include <assert.h>
#include <string.h>
#include <stdlib.h>

#include "net.h"
#include "log.h"
#include "socket_udp.h"
#include "socket_tcp.h"
#include "poll.h"
#include "utils.h"
#include "sys.h"

#if defined MN_APPLE  || defined MN_ANDROID
#include <sys/time.h>
#endif

#if defined  MN_APPLE || defined MN_ANDROID
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <sys/types.h>
#endif


int parse_url(const char *url,struct mn_sockaddr *addr)
{
    char *colon = NULL;
    if (NULL == url || NULL == addr)
    {
        return MN__EARG;
    }
    memset(addr->host,'\0',MAX_HOST_LEN);
    addr->port = 0;
    addr->proto = NET_UNKNOWN;
    
    if (! memcmp((void *)url,"udp://",6)) {
        addr->proto = NET_UDP;
        colon =  strchr(url+6,':');
        if (!colon) {
            return MN__EURL;
        }
        
        ssize_t ip_len = colon - url-6 ;
        if (ip_len <= 0 || (6+ip_len) >= strlen(url)) {
            return MN__EURL;
        }
        memcpy(addr->host,url+6,ip_len);
        
    } else if (! memcmp((void *)url,"tcp://",6)) {
        addr->proto = NET_TCP;
        colon =  strchr(url+6,':');
        if (!colon) {
            return MN__EURL;
        }
        
        ssize_t ip_len = colon - url-6 ;
        if (ip_len <= 0 || (6+ip_len) >= strlen(url)) {
            return MN__EURL;
        }
        memcpy(addr->host,url+6,ip_len);
        
    } else {
        addr->proto = NET_UNKNOWN;
        return MN__EURL;
    }
    
    int port = atoi(colon+1);
    if (port<=0 || port >65536) {
        return MN__EURL;
    }
    addr->port = port;
    return 0;
}


static int connect_timeout(const struct mn_socket *socket, uint64_t timeout)
{
    if(NULL == socket || socket->sfd<3 ||  timeout ==0) {
        return MN__EARG;
    }
    
    struct timeval bt;
    gettimeofday(&bt, NULL);
    int err;
    err = connect(socket->sfd, &socket->dest_addr, socket->addrlen);
    
    if (0 == err) {
        return 0;
    }
    
    // use poll to instead select
    if (errno == EINPROGRESS){
        //FD_SETSIZE check
        
        if (socket->sfd >= FD_SETSIZE) {
            return MN__ECONN;
        }
        
        // select here
        fd_set rfds,wfds;
        struct timeval tv;
        int isConn = 0;
        while (!isConn){
            
            FD_ZERO(&rfds);
            FD_ZERO(&wfds);
            FD_SET(socket->sfd,&rfds);
            FD_SET(socket->sfd,&wfds);
            
            tv.tv_sec  = timeout/1000;
            tv.tv_usec = (timeout%1000) * 1000;
            
            int rst = select (socket->sfd+1, &rfds, &wfds, NULL, &tv);
            struct timeval st;
            gettimeofday(&st, NULL);
            if (timeval_min_usec(&st, &bt) > timeout) {
                return MN__ETIMEOUT;
            }
            if (rst<0){
                //select error
                LOG_E("select error \n");
                return MN__ECONN;
            } else if (rst == 0) {
                // Timeout :haven't done connection
                continue;
            } else {
                if (FD_ISSET(socket->sfd, &wfds) && !(FD_ISSET(socket->sfd, &rfds))) {
                    // can write but not readable;ok connected
                    isConn = 1;
                    continue;
                } else if (FD_ISSET(socket->sfd, &wfds) && FD_ISSET(socket->sfd, &rfds)) {
                    int err = connect(socket->sfd, &socket->dest_addr, socket->addrlen);
                    if (err) {
                        if (errno == EISCONN) {
                            // connected
                            isConn = 1;
                            continue;
                        } else {
                            continue;
                        }
                    } else {
                        isConn = 1;
                        continue;
                    }
                } else {
                    // unknow fd status
                    continue;
                }
            }
            
        }  //end of while
    } else {
        LOG_E("errno is %d",errno);
        return MN__ECONN;
    } // May have other errno s;
    
    return 0;
}

int mn_net_close(struct mn_socket *sfd)
{
    if (NULL == sfd) {
        return MN__EARG;
    }
    
    int rst;

    rst = mn_socket_close(sfd);
    return rst;
}

int mn_net_connect(struct mn_socket *sfd, const char *url, uint64_t timeout)
{
    struct mn_sockaddr addr;
    if (NULL == url || NULL == sfd) {
		return MN__ENULL;
    }
    // parse url
    int rst  = parse_url(url,&addr);
    if (0 != rst) {
        return rst;
    }
    // create socket
    if (NET_UDP == addr.proto) {
        rst = mn_socket_udp(&addr, sfd);
    } else if (NET_TCP == addr.proto) {
        rst = mn_socket_tcp(&addr, sfd);
    } else {
        rst =  MN__EPROTO;
    }
    
    if (0 != rst) {
        return rst;
    }
        
    // set nonblock
    mn_socket_setnonblock(sfd);
    
    // set send & recv buf
    mn_socket_setsendbuff(sfd, NET_SEND_BUF_SIZE);
    mn_socket_setrecvbuff(sfd, NET_RECV_BUF_SIZE);
    
    // set no delay
    mn_socket_setnodelay(sfd);
    
    // ignore pipe siganl
    mn_sys_ignore_pipe();
    
    // do connect for tcp
    if (NET_TCP == addr.proto) {
        int ret = connect_timeout(sfd, timeout);
        if (ret) {
            mn_socket_close(sfd);
            return ret;
        }
    }
    
    return 0;
}

int mn_net_send(struct mn_socket *sfd,const void *buf,size_t *len,uint64_t timeout)
{
    int rst;
	
    if (NULL == sfd || NULL == buf || NULL == len) {
		return MN__EARG;
    }
	
    if (NET_TCP ==  sfd->proto) {
#if defined MN_ANDROID
        rst = mn_socket_send(sfd, buf, len, MSG_NOSIGNAL, timeout);
#else
        rst = mn_socket_send(sfd, buf, len, 0, timeout);
#endif
    } else  if (NET_UDP) {
        ssize_t ret = mn_socket_sendto(sfd, buf, *len, 0, timeout);
        *len = ret;
        if (ret > 0) {
            rst  = 0;
        } else {
            rst = (int)ret;
        }
    } else {
        rst = MN__EPROTO;
    }
    
    return rst;
}

int mn_net_recv(struct mn_socket *sfd,void *buf,size_t *len,uint64_t timeout)
{
    int rst;
    if (NULL == sfd || NULL == buf || NULL == len) {
		return MN__EARG;
    }
    
    if (NET_TCP == sfd->proto) {
        rst = mn_socket_recv(sfd, buf, len, 0, timeout);
    } else if (NET_UDP == sfd->proto) {
        ssize_t ret = mn_socket_recvfrom(sfd, buf, *len, 0, timeout);
        *len = ret;
        if (ret >0) {
            rst = 0;
        } else {
            rst = (int) ret;
        }
    } else {
        rst = 0;
    }
    
    return rst;
}
    

    



