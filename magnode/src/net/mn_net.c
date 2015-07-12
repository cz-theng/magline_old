//
//  mn_net.cpp
//  magnode
//
//  Created by cz on 15/7/7.
//  Copyright (c) 2015年cz. All rights reserved.
//

#include <assert.h>
#include <string.h>
#include <stdlib.h>

#include "mn_net.h"
#include "mn_log.h"
#include "mn_socket_udp.h"
#include "mn_socket_tcp.h"
#include "mn_poll.h"


int parse_url(const char *url,struct net_sockaddr *addr)
{
    assert(url);
    assert(addr);
    memset(addr->host,'\0',HOST_LEN);
    addr->port = 0;
    
    if (! memcmp((void *)url,"udp",3)) {
        addr->proto = NET_UDP;
    } else if (! memcmp((void *)url,"tcp",3)) {
        addr->proto = NET_TCP;
    } else {
        addr->proto = NET_UNKNOWN;
    }
    
    char *colon = strchr(url+5,':');
    if (!colon) {
        return MN_EURL;
    }
    
    ssize_t ip_len = colon - url-6 ;
    memcpy(addr->host,url+6,ip_len);
    
    int port = atoi(colon+1);
    if (port<=0) {
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
    rst = mn_socket_close(sfd);
    return rst;
}

int mn_connect(const char *url,struct mn_socket *sfd, uint64_t timeout)
{

    struct net_sockaddr addr;
    // parse url
    int rst  = parse_url(url,&addr);
    if (rst <0) {
        return rst;
    }
    // create socket
    if (NET_UDP == addr.proto) {
        rst = mn_socket_udp(&addr, sfd);
    } else if (NET_TCP == addr.proto) {
        rst = -1;
    } else {
        rst =  MN_EPROTO;
    }
    
    if (!rst) {
        return rst;
    }
        
    // set socketopt
    mn_socket_setsocketopt_nonblock(sfd);
    
    // do connect for tcp
    
    return -1;
}

ssize_t mn_send(struct mn_socket *sfd,const void *buf,size_t len,uint64_t timeout)
{
    ssize_t rst;
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
    int ret;
    ssize_t rst;
    
    ret = mn_poll(sfd->sfd, MN_POLL_OUT,timeout);
    
    if (ret < 0) {
        if (ret == -1) {
            // time out
            return -2;
        } else {
            mn_log("poll error !");
        }
        return ret;
    }
    
    if (NET_TCP == sfd->proto) {
        rst = mn_socket_recv(sfd,buf,len,0);
    } else if (NET_UDP == sfd->proto) {
        rst = mn_socket_recvfrom(sfd,buf,len,0);
    } else {
        mn_log("Unknown proto!\n");
        rst = 0;
    }
    
    return rst;
}
    

    

    

