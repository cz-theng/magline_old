/**
 * Author :cz cz.theng@gmail.com
 * Licence MIT
 */

#include "mn_socket_udp.h"
#include <errno.h>
#include <stdio.h>


ssize_t mn_socket_sendto(struct mn_socket *fd, const void *buf, size_t len, int flags)
{
    ssize_t rst = 0;
	if (NULL == fd || NULL == buf)
		return 0;
    rst = sendto(fd->sfd,(const char *) buf, len, flags, &fd->dest_addr, fd->addrlen);
    if (rst < 0) {
        return  MN_ESEND;
    } else {
        return rst;
    }
}

ssize_t mn_socket_recvfrom(struct mn_socket *fd, void *buf, size_t len, int flags)
{
    ssize_t rst = 0;
	if (NULL == fd || NULL == buf)
		return 0;
    rst = recvfrom(fd->sfd,(char *) buf, len, flags,  NULL, NULL);
    if (rst >= 0) {
        return rst;
    } else {
        return MN_ERECVFROM;
    }
        
}

ssize_t mn_socket_sendmsg(struct mn_socket *fd, const struct msghdr *msg, int flags)
{
    return -1;
}

ssize_t mn_socket_recvmsg(struct mn_socket *fd, struct msghdr *msg, int flags)
{

    
    return -1;
}

