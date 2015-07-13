//
//  mn_socket_udp.h
//  magnode
//
//  Created by cz on 15/7/7.
//  Copyright (c) 2015年cz. All rights reserved.
//

#ifndef __magnode__mn_socket_udp__
#define __magnode__mn_socket_udp__

#include "mn_socket.h"


ssize_t mn_socket_sendto(struct mn_socket *fd, const void *buf, size_t len, int flags);

ssize_t mn_socket_recvfrom(struct mn_socket *fd, void *buf, size_t len, int flags);

ssize_t mn_socket_sendmsg(struct mn_socket *fd, const struct msghdr *msg, int flags);

ssize_t mn_socket_recvmsg(struct mn_socket *fd, struct msghdr *msg, int flags);
    


#endif /* defined(__magnode__mn_socket_udp__) */
