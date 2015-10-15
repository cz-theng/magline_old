/**
 * Author :cz cz.theng@gmail.com
 * Licence MIT
 */

#ifndef MAGNODE_NET_MN_SOCKET_UDP_H_
#define MAGNODE_NET_MN_SOCKET_UDP_H_

#include "socket.h"


ssize_t mn_socket_sendto(struct mn_socket *fd, const void *buf, size_t len, int flags, uint64_t timeout);

ssize_t mn_socket_recvfrom(struct mn_socket *fd, void *buf, size_t len, int flags, uint64_t timeout);

ssize_t mn_socket_sendmsg(struct mn_socket *fd, const struct msghdr *msg, int flags, uint64_t timeout);

ssize_t mn_socket_recvmsg(struct mn_socket *fd, struct msghdr *msg, int flags, uint64_t timeout);
    


#endif /* defined(MAGNODE_NET_MN_SOCKET_UDP_H_) */
