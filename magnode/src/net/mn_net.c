/**
 * Author :cz cz.theng@gmail.com
 * Licence MIT
 */

#include <assert.h>
#include <string.h>
#include <stdlib.h>

#include "mn_net.h"
#include "mn_log.h"
#include "mn_socket_udp.h"
#include "mn_socket_tcp.h"
#include "mn_poll.h"
#include "mn_utils.h"

#if defined MN_APPLE  || defined MN_ANDROID
#include <sys/time.h>
#endif


int parse_url(const char *url,struct mn_sockaddr *addr)
{
    char *colon = NULL;
    if (NULL == url || NULL == addr)
    {
        return MN_EARG;
    }
    memset(addr->host,'\0',MAX_HOST_LEN);
    addr->port = 0;
    addr->proto = NET_UNKNOWN;
    
    if (! memcmp((void *)url,"udp://",6)) {
        addr->proto = NET_UDP;
        colon =  strchr(url+6,':');
        if (!colon) {
            return MN_EURL;
        }
        
        ssize_t ip_len = colon - url-6 ;
        if (ip_len <= 0 || (6+ip_len) >= strlen(url)) {
            return MN_EURL;
        }
        memcpy(addr->host,url+6,ip_len);
        
    } else if (! memcmp((void *)url,"tcp://",6)) {
        addr->proto = NET_TCP;
        colon =  strchr(url+6,':');
        if (!colon) {
            return MN_EURL;
        }
        
        ssize_t ip_len = colon - url-6 ;
        if (ip_len <= 0 || (6+ip_len) >= strlen(url)) {
            return MN_EURL;
        }
        memcpy(addr->host,url+6,ip_len);
        
    } else {
        addr->proto = NET_UNKNOWN;
        return MN_EURL;
    }
    
    int port = atoi(colon+1);
    if (port<=0 || port >65536) {
        return MN_EURL;
    }
    addr->port = port;
    return 0;
}


static int connect_timeout(const struct mn_socket *socket, uint64_t timeout)
{
    if(NULL == socket || socket->sfd<3 ||  timeout ==0) {
        return MN_EARG;
    }
    
    struct timeval bt;
    gettimeofday(&bt, NULL);
    int err;
    err = connect(socket->sfd, &socket->dest_addr, socket->addrlen);
    
    if (0 == err) {
        return 0;
    }
    
    if (errno == EINPROGRESS){
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
                return MN_ETIMEOUT;
            }
            if (rst<0){
                //select error
                LOG_E("select error \n");
                return MN_ECONN;
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
        return MN_ECONN;
    } // May have other errno s;
    
    return 0;
}


int mn_listen(char *url)
{
    return 0;
}

int mn_close(struct mn_socket *sfd)
{
    if (NULL == sfd) {
        return MN_EARG;
    }
    
    int rst;

    rst = mn_socket_close(sfd);
    return rst;
}

int mn_connect(const char *url,struct mn_socket *sfd, uint64_t timeout)
{
    struct mn_sockaddr addr;
    if (NULL == url || NULL == sfd) {
		return MN_ENULL;
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
        rst =  MN_EPROTO;
    }
    
    if (0 != rst) {
        return rst;
    }
        
    // set socketopt
    mn_socket_setnonblock(sfd);
    
    // set send & recv buf
    mn_socket_setsendbuff(sfd, NET_SEND_BUF_SIZE);
    mn_socket_setrecvbuff(sfd, NET_RECV_BUF_SIZE);
    
    // do connect for tcp
    if (NET_TCP == addr.proto) {
        int ret = connect_timeout(sfd, timeout);
        if (ret) {
            return ret;
        }
    }
    
    return 0;
}

int mn_send(struct mn_socket *sfd,const void *buf,size_t *len,uint64_t timeout)
{
    int rst;
	
    if (NULL == sfd || NULL == buf || NULL == len) {
		return MN_EARG;
    }
	
    if (NET_TCP ==  sfd->proto) {
        rst = mn_socket_send(sfd, buf, len, 0, timeout);
    } else  if (NET_UDP) {
        ssize_t ret = mn_socket_sendto(sfd, buf, *len, 0, timeout);
        *len = ret;
        if (ret > 0) {
            rst  = 0;
        } else {
            rst = (int)ret;
        }
    } else {
        rst = MN_EPROTO;
    }
    
    return rst;
}

int mn_recv(struct mn_socket *sfd,void *buf,size_t *len,uint64_t timeout)
{
    int rst;
    if (NULL == sfd || NULL == buf || NULL == len) {
		return MN_EARG;
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
    

    



