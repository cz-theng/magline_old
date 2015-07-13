//
//  mn_socket.cpp
//  magnode
//
//  Created by cz on 15/7/7.
//  Copyright (c) 2015年cz. All rights reserved.
//

#include "mn_socket.h"

#include <stddef.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <stdlib.h>
#include <sys/fcntl.h>

int mn_socket_udp(struct mn_sockaddr *addr, struct mn_socket *sfd)
{
    if (NET_UDP != addr->proto) {
        return MN_EPROTO;
    } else {
        sfd->proto = addr->proto;
        sfd->sfd = socket(AF_INET, SOCK_DGRAM, 0);
        
        struct sockaddr_in *udpaddr =(struct sockaddr_in*) &sfd->dest_addr;
        int rst = inet_aton(addr->host, &udpaddr->sin_addr);
        // TODO: adjust rst;
    
        udpaddr->sin_family = AF_INET;
        udpaddr->sin_port = htons(addr->port);
        
        sfd->addrlen = sizeof(struct sockaddr_in);
        return 0;
    }
}

int mn_socket_tcp(struct mn_sockaddr *addr, struct mn_socket *sfd)
{
    return -1;
}

int mn_socket_close(struct mn_socket *sfd)
{
    if (!sfd) {
        return MN_ENULLARG;
    }
    close(sfd->sfd);
    return 0;
}

int mn_socket_setsocketopt_nonblock(struct mn_socket *sfd)
{
    int flags;
    
    flags = fcntl(sfd->sfd, F_GETFL, 0);
    
    flags	|=	O_NONBLOCK | O_ASYNC;

    
    return fcntl(sfd->sfd, F_SETFL, flags);
}


