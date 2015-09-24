/**
 * Author :cz cz.theng@gmail.com
 * Licence MIT
 */

#include <assert.h>
#include <string.h>
#include <stdlib.h>

#include "mn_net.h"
#include "mn_socket_udp.h"
#include "mn_socket_tcp.h"
#include "mn_poll.h"

static int parse_url(const char *url,struct mn_sockaddr *addr)
{
    char *colon = NULL;
	if (NULL == url || NULL == addr)
	{
		return MN_ENULL;
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


int mn_listen(char *url)
{
    return 0;
}

int mn_close(struct mn_socket *sfd)
{
    int rst;

    if (NULL == sfd) {
		return 0;
    }
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
    
    }
    
    return 0;
}

ssize_t mn_send(struct mn_socket *sfd,const void *buf,size_t len,uint64_t timeout)
{
    ssize_t rst;
	
	if (NULL == sfd || NULL == buf)
		return 0;
	
    if (NET_TCP ==  sfd->proto) {
        rst = mn_socket_send(sfd, buf, len, 0);
    } else  if (NET_UDP) {
        rst = mn_socket_sendto(sfd, buf, len, 0);
    } else {
        rst = MN_EPROTO;
    }
    
    return rst;
}

ssize_t mn_recv(struct mn_socket *sfd,void *buf,size_t len,uint64_t timeout)
{
    int ret = 0;
    ssize_t rst;
    if (NULL == sfd || NULL == buf)
		return 0;
    ret = mn_poll(sfd->sfd, MN_POLL_OUT,timeout);
    if (ret < 0) {
        return ret;
    }
    
    if (NET_TCP == sfd->proto) {
        rst = mn_socket_recv(sfd,buf,len,0);
    } else if (NET_UDP == sfd->proto) {
        rst = mn_socket_recvfrom(sfd,buf,len,0);
    } else {
        rst = 0;
    }
    
    return rst;
}
    

    
#ifdef MN_ANDROID
uint64_t htonll(uint64_t val) {
    return (((uint64_t) htonl(val)) << 32) + htonl(val >> 32);
}

uint64_t ntohll(uint64_t val) {
    return (((uint64_t) ntohl(val)) << 32) + ntohl(val >> 32);
}
#endif


