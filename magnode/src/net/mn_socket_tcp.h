//
//  mn_socket_tcp.h
//  magnode
//
//  Created by cz on 15/7/7.
//  Copyright (c) 2015年cz. All rights reserved.
//

#ifndef __magnode__mn_socket_tcp__
#define __magnode__mn_socket_tcp__


#include "mn_socket.h"
#include "mn_socket_tcp.h"


size_t mn_socket_send(struct mn_socket *fd, const void *buf, size_t len, int flags);

ssize_t mn_socket_recv(struct mn_socket *fd, void *buf, size_t len, int flags);


#endif /* defined(__magnode__mn_socket_tcp__) */
