/**
 * Author :cz cz.theng@gmail.com
 * Licence MIT
 */

#include "socket_tcp.h"
#include "poll.h"
#include "utils.h"

#if defined MN_APPLE  || defined MN_ANDROID
#include <sys/time.h>
#endif

// send or write
ssize_t	socket_send(int fd, const void *buf, size_t len, int flag)
{
    ssize_t rst  = 0;
    rst = send(fd, buf, len, flag);
    if (rst > 0) {
        return rst;
    } else if (0 == rst) {
        return MN__ECLOSED;
    } else if (rst <0 &&(errno==EINTR || errno==EAGAIN || errno==EWOULDBLOCK)) {
        return 0;
    } else {
        return rst;
    }
    return rst;
}

// recv or read
ssize_t	socket_recv(int fd, void *buf, size_t len, int flag)
{
    ssize_t rst = 0;
    rst = recv(fd, buf, len, flag);
    if (rst > 0) {
        return rst;
    } else if (0 == rst) {
        return MN__ECLOSED;
    } else if (rst <0 &&(errno==EINTR || errno==EAGAIN || errno==EWOULDBLOCK)) {
        return 0;
    } else {
        return rst;
    }
    return  rst;
}

int mn_socket_send(struct mn_socket *fd, const void *buf, size_t *len, int flags, uint64_t timeout)
{
    if (NULL == fd || buf == fd || NULL==len) {
        return MN__EARG;
    }
    size_t task;
    const char *guard;
    
    struct timeval bt;
    gettimeofday(&bt, NULL);
    
    guard=buf;
    task=*len;
    int pollrst = 0;
    while (task >0) {
        struct timeval et;
        gettimeofday(&et, NULL);
        long rtimeout = timeout - timeval_min_usec(&et, &bt);
        rtimeout = rtimeout > 0 ? rtimeout : 0;
        if (0 == timeout) {
            rtimeout = 0;
        }
        pollrst = mn_poll(fd->sfd, MN_POLL_OUT, rtimeout);
        if (pollrst !=0 ) {
            if (MN__ETIMEOUT == pollrst) {
                return MN__ETIMEOUT;
            } else {
                return MN__EPOLL;
            }
        }
        
        ssize_t rst = socket_send(fd->sfd, guard, task, flags);
        if (0 == timeout) {
            if (rst > 0) {
                *len = rst;
                return 0;
            } else if (rst == 0) {//errno==EINTR || errno==EAGAIN || errno==EWOULDBLOCK
                *len= 0;
                return 0;
            } else if (MN__ECLOSED == rst){
                *len = 0;
                return MN__ECLOSED;
            } else {
                *len = 0;
                return MN__ESEND;
            }
        }
        
        if (rst > 0) {
            task -= rst;
            guard += rst;
        } else if (0 == rst) { //errno==EINTR || errno==EAGAIN || errno==EWOULDBLOCK
            continue;
        } else if (MN__ECLOSED == rst) {
            //close
            *len = (*len) - task;
            return MN__ECLOSED;
        } else { // ret<0
            *len = (*len) - task;
            // other errors
            return MN__ESEND;
        }
    }

    return 0;
}

int mn_socket_recv(struct mn_socket *fd, void *buf, size_t *len, int flags, uint64_t timeout)
{
    if (NULL == fd || NULL == buf || NULL == len) {
        return MN__EARG;
    }
    
    size_t task;
    char *guard;
    
    struct timeval bt;
    gettimeofday(&bt, NULL);
    
    guard = buf;
    task = *len;
    int pollrst = 0;
    while (task > 0) {
        struct timeval et;
        gettimeofday(&et, NULL);
        long rtimeout = timeout - timeval_min_usec(&et, &bt);
        rtimeout = rtimeout> 0 ? rtimeout : 0;
        if (0 == timeout) {
            rtimeout = 0;
        }
        pollrst = mn_poll(fd->sfd, MN_POLL_IN, rtimeout);
        if (pollrst !=0 ) {
            if (MN__ETIMEOUT == pollrst) {
                return MN__ETIMEOUT;
            } else {
                return MN__EPOLL;
            }
        }
        
        ssize_t rst = socket_recv(fd->sfd, guard, task, flags);
        if (0 == timeout) {
            if (rst > 0) {
                *len = rst;
                return 0;
            } else if (rst == 0) { //errno==EINTR || errno==EAGAIN || errno==EWOULDBLOCK
                *len = 0;
                return 0;
            } else if (MN__ECLOSED == rst) {
                *len = 0;
                return MN__ECLOSED;
            } else {
                *len = 0;
                return MN__ERECV;
            }
        }
        
        if (rst > 0) {
            guard += rst;
            task -= rst;
        } else if (0 == rst) { //errno==EINTR || errno==EAGAIN || errno==EWOULDBLOCK
            continue;
        } else if (MN__ECLOSED == rst) {
            //close
            *len = (*len) - task;
            return MN__ECLOSED;
        } else {
            // other errors
            *len = (*len) - task;
            return MN__ESEND;
        }
    }
    return 0;
}

