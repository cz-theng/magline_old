/**
 * Author :cz cz.theng@gmail.com
 * Licence MIT
 */

#ifndef MAGNODE_NET_MN_SOCKET_TCP_H_
#define MAGNODE_NET_MN_SOCKET_TCP_H_


#include "mn_socket.h"
#include "mn_socket_tcp.h"


size_t mn_socket_send(struct mn_socket *fd, const void *buf, size_t len, int flags);

ssize_t mn_socket_recv(struct mn_socket *fd, void *buf, size_t len, int flags);


#endif /* defined(MAGNODE_NET_MN_SOCKET_TCP_H_) */
