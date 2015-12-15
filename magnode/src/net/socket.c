/**
 * Author :cz cz.theng@gmail.com
 * Licence MIT
 */

#include <ctype.h>
#include <stddef.h>
#include <string.h>
#include <stdlib.h>
#include <fcntl.h>
#include <errno.h>

#include "socket.h"

#if defined  MN_APPLE || defined MN_ANDROID
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#endif

#if defined  MN_WIN
#include <winsock2.h>
#include <WS2tcpip.h>
#include <windows.h>


# if defined(_MSC_VER)
#  include <BaseTsd.h>
typedef SSIZE_T ssize_t;
//typedef SIZE_T  size_t;
# endif

int inet_aton(const char* cp, struct in_addr* inp)
{
	unsigned long ulAddr;

	ulAddr = inet_addr(cp);

	if (INADDR_NONE == ulAddr) {
		inp->S_un.S_addr	=	ulAddr;
		return 0;
	}

	inp->S_un.S_addr	=	ulAddr;

	return 1;
}
#endif





static int ip_or_domain(const char *host)
{
    if (! host) {
        return -1;
    }
    int i;
    for (i=0;i< strlen(host); i++) {
        if ( (host[i]<='9' && host[i]>='0') || '.' == host[i]) {
            continue;
        } else {
            return 1;
        }
    }
    
    return 0;
}

static int host2in(const char *host, struct in_addr *addr)
{
    if (! host || !addr ){
        return MN__ENULL;
    }
    int isIP  = ip_or_domain(host);
    if (0==isIP) {
        inet_aton(host, addr);
    } else if (1 == isIP){
        struct hostent* hosts = gethostbyname(host);
        if (!hosts) {
            return -1;
        }
        memcpy(addr, hosts->h_addr_list[0], sizeof(struct in_addr));
    } else {
        return -1;
    }
    
    return 0;
}

int mn_socket_udp(struct mn_sockaddr *addr, struct mn_socket *sfd)
{
	if (NULL == addr || NULL == sfd)
		return MN__ENULL;
    if (NET_UDP != addr->proto) {
        return MN__EPROTO;
    } else {
        sfd->proto = addr->proto;
        sfd->sfd = socket(AF_INET, SOCK_DGRAM, 0);
        
        struct sockaddr_in *udpaddr =(struct sockaddr_in*) &sfd->dest_addr;
        //int rst = inet_aton(addr->host, &udpaddr->sin_addr);
        int rst = host2in(addr->host, &udpaddr->sin_addr);
        if (0 != rst) {
            return -1;
        }
    
        udpaddr->sin_family = AF_INET;
        udpaddr->sin_port = htons(addr->port);
        
        sfd->addrlen = sizeof(struct sockaddr_in);
        return 0;
    }
}

int mn_socket_tcp(struct mn_sockaddr *addr, struct mn_socket *sfd)
{
    if (NULL == addr || NULL == sfd)
        return MN__ENULL;
    if (NET_TCP != addr->proto) {
        return MN__EPROTO;
    } else {
        sfd->proto = addr->proto;
        sfd->sfd = socket(AF_INET, SOCK_STREAM, 0);
        
        struct sockaddr_in *tcpaddr =(struct sockaddr_in*) &sfd->dest_addr;
        int rst = host2in(addr->host, &tcpaddr->sin_addr);
        if (0 != rst) {
            return MN__EHOST;
        }
        
        tcpaddr->sin_family = AF_INET;
        tcpaddr->sin_port = htons(addr->port);
        
        sfd->addrlen = sizeof(struct sockaddr_in);
        return 0;
    }
}

int mn_socket_close(struct mn_socket *sfd)
{
    if (!sfd) {
        return MN__ENULLARG;
    }
#if defined MN_WIN
	return closesocket(sfd->sfd);
#else
    if (sfd->sfd <3) {
        return 0;
    }
    close(sfd->sfd);
    return 0;
#endif
    return 0;
}

int mn_socket_setnonblock(struct mn_socket *sfd)
{
    int flags;

	if (NULL == sfd)
		return -1;
    
#if defined MN_WIN
	unsigned long lParam =0;

	return ioctlsocket(sfd->sfd, FIONBIO, &lParam);

#endif
#if defined  MN_APPLE || defined MN_ANDROID
    flags = fcntl(sfd->sfd, F_GETFL, 0);
    
    flags	|=	O_NONBLOCK | O_ASYNC;
    
    return fcntl(sfd->sfd, F_SETFL, flags);
#endif
}

int mn_socket_setrecvbuff(struct mn_socket *sfd, int size)
{
    return setsockopt(sfd->sfd, SOL_SOCKET, SO_RCVBUF, (void*)&size, sizeof(size));
}

int mn_socket_setsendbuff(struct mn_socket *sfd, int size)
{
    return setsockopt(sfd->sfd, SOL_SOCKET, SO_SNDBUF, (void*)&size, sizeof(size));
}

int mn_socket_setnodelay(struct mn_socket *sfd)
{
    int flag = 1;
    return setsockopt(sfd->sfd, IPPROTO_TCP, TCP_NODELAY, (const char *)&flag, sizeof(flag));
}


