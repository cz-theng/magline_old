/**
 * Author :cz cz.theng@gmail.com
 * Licence MIT
 */

#ifndef MAGNODE_NET_MN_NET_H_
#define MAGNODE_NET_MN_NET_H_

#include "magnode.h"
#include "os.h"

#if defined MN_APPLE || defined MN_ANDROID
#include <sys/select.h>
#include <sys/socket.h>
#include <sys/un.h>
#include <netdb.h>
#include <unistd.h>
#include <netinet/in.h>
#include <netinet/tcp.h>
#include <arpa/inet.h>
#include <sys/socket.h>
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
#endif

#include <stdint.h>
#include <stddef.h>


#ifdef __cplusplus

extern "C" {
#endif

    #define MAX_HOST_LEN 256
        
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
    #ifndef MN_ENULL
    #define MN_ENULL        -(MN_HAUSNUMERO + 8)
    #endif
    #ifndef MN_ECONN
    #define MN_ECONN        -(MN_HAUSNUMERO + 9)
    #endif
    #ifndef MN_EARG
    #define MN_EARG         -(MN_HAUSNUMERO + 10)
    #endif
    #ifndef MN_ESENDTO
    #define MN_ESENDTO      -(MN_HAUSNUMERO + 11)
    #endif
    #ifndef MN_ERECV
    #define MN_ERECV        -(MN_HAUSNUMERO + 12)
    #endif
    
    
    enum net_proto
    {
        NET_TCP,
        NET_UDP,
        NET_UNKNOWN,
    };
    
    enum net_bufsize
    {
        NET_RECV_BUF_SIZE = 1*1024*1024,
        NET_SEND_BUF_SIZE = 1*1024*1024,
    };
/*
    struct mn_socket
    {
        int sfd;
        enum net_proto proto;
        struct sockaddr dest_addr;
        socklen_t addrlen;
    };
*/
    struct mn_sockaddr
    {
        enum net_proto proto;
        char host[MAX_HOST_LEN];
        uint16_t port;
    };

    int mn_net_listen(char *url);

    int mn_net_close(struct mn_socket *sfd);

    int mn_net_connect(const char *url,struct mn_socket *sfd, uint64_t timeout);

    int mn_net_send(struct mn_socket *sfd,const void *buf,size_t *len,uint64_t timeout);

    int mn_net_recv(struct mn_socket *sfd,void *buf,size_t *len,uint64_t timeout);
    
            
#ifdef __cplusplus
}
#endif

#endif /* defined(MAGNODE_NET_MN_NET_H_) */
