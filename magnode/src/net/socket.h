/**
 * Author :cz cz.theng@gmail.com
 * Licence MIT
 */

#ifndef MAGNODE_NET_MN_SOCKET_H_
#define MAGNODE_NET_MN_SOCKET_H_

#include "net.h"

int mn_socket_udp(struct mn_sockaddr *addr, struct mn_socket *sfd);

int mn_socket_tcp(struct mn_sockaddr *addr, struct mn_socket *sfd);

int mn_socket_close(struct mn_socket *sfd);

int mn_socket_setnonblock(struct mn_socket *sfd);

int mn_socket_setrecvbuff(struct mn_socket *sfd, int size);

int mn_socket_setsendbuff(struct mn_socket *sfd, int size);

int mn_socket_setnodelay(struct mn_socket *sfd);

#endif /* defined(MAGNODE_NET_MN_SOCKET_H_) */
