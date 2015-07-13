//
//  mn_socket.h
//  magnode
//
//  Created by cz on 15/7/7.
//  Copyright (c) 2015年cz. All rights reserved.
//

#ifndef __magnode__mn_socket__
#define __magnode__mn_socket__

#include "mn_net.h"

int mn_socket_udp(struct mn_sockaddr *addr, struct mn_socket *sfd);

int mn_socket_tcp(struct mn_sockaddr *addr, struct mn_socket *sfd);

int mn_socket_close(struct mn_socket *sfd);

int mn_socket_setsocketopt_nonblock(struct mn_socket *sfd);


#endif /* defined(__magnode__mn_socket__) */
