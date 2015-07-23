//
//  mn_net.h
//  magnode
//
//  Created by cz on 15/7/7.
//  Copyright (c) 2015年cz. All rights reserved.
//

#ifndef __magnode__mn_net__
#define __magnode__mn_net__

#include <sys/select.h>
#include <sys/socket.h>
#include <sys/un.h>
#include <netdb.h>
#include <unistd.h>
#include <netinet/in.h>
#include <netinet/tcp.h>
#include <arpa/inet.h>
#include <sys/socket.h>

#include <stdint.h>
#include <stddef.h>


#ifdef __cplusplus

extern "C" {
#endif

    #define HOST_LEN 16
        
    #define MN_HAUSNUMERO 520727134
    
    #ifndef MN_ETIMEOUT
    #define MN_ETIMEOUT     -(MN_HAUSNUMERO + 1)
    #endif
    #ifndef MN_EPROTO
    #define MN_EPROTO       -(MN_HAUSNUMERO + 2)
    #endif
    #ifndef MN_ESEND
    #define MN_ESEND        -(MN_HAUSNUMERO + 3)
    #endif
    #ifndef MN_ERECVFROM
    #define MN_ERECVFROM    -(MN_HAUSNUMERO + 4)
    #endif
    #ifndef MN_ENULLARG
    #define MN_ENULLARG     -(MN_HAUSNUMERO + 5)
    #endif
    #ifndef MN_EURL
    #define MN_EURL         -(MN_HAUSNUMERO + 6)
    #endif
    #ifndef MN_EPOLL
    #define MN_EPOLL        -(MN_HAUSNUMERO + 7)
    #endif
    
    
    enum net_proto
    {
        NET_TCP,
        NET_UDP,
        NET_UNKNOWN,
    };
    
    struct mn_socket
    {
        int sfd;
        enum net_proto proto;
        struct sockaddr dest_addr;
        socklen_t addrlen;
    };
    
    struct mn_sockaddr
    {
        enum net_proto proto;
        char host[HOST_LEN];
        uint16_t port;
    };

    int mn_listen(char *url);

    int mn_close(struct mn_socket *sfd);

    int mn_connect(const char *url,struct mn_socket *sfd, uint64_t timeout);

    ssize_t mn_send(struct mn_socket *sfd,const void *buf,size_t len,uint64_t timeout);

    ssize_t mn_recv(struct mn_socket *sfd,void *buf,size_t len,uint64_t timeout);
#ifdef __cplusplus
}
#endif

#endif /* defined(__magnode__mn_net__) */
