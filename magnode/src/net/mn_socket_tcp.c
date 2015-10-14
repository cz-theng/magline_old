/**
 * Author :cz cz.theng@gmail.com
 * Licence MIT
 */

#include "mn_socket_tcp.h"
#include "mn_poll.h"
#include "mn_utils.h"

#if defined MN_APPLE  || defined MN_ANDROID
#include <sys/time.h>
#endif

int mn_socket_send(struct mn_socket *fd, const void *buf, size_t *len, int flags, uint64_t timeout)
{
    if (NULL == fd || buf == fd || NULL==len) {
        return MN_EARG;
    }
    size_t task;
    const char *guard;
    
    struct timeval bt;
    gettimeofday(&bt, NULL);
    
    guard=buf;
    task=*len;
    while (task >0) {
        ssize_t rst = send(fd->sfd, guard, task, flags);
        if (0 == timeout) {
            if (rst > 0) {
                *len = rst;
                return 0;
            } else {
                *len = 0;
                return MN_ESEND;
            }
        }
        
        struct timeval st;
        gettimeofday(&st, NULL);
        if (timeval_min_usec(&st, &bt) > timeout) {
            if (rst >0) {
                *len = (*len) - (task - rst);
            } else {
                *len = (*len) - task;
            }
            return MN_ETIMEOUT;
        }
        if (rst > 0) {
            task -= rst;
            guard += rst;
        } else if (0 == rst) {
            //close
            *len = (*len) - task;
            return 0;
        } else { // ret<0
            if (errno==EINTR || errno==EAGAIN) {
                continue;
            } else {
                *len = (*len) - task;
                // other errors
                return MN_ESEND;
            }
        }
    }

    return 0;
}

int mn_socket_recv(struct mn_socket *fd, void *buf, size_t *len, int flags, uint64_t timeout)
{
    if (NULL == fd || NULL == buf || NULL == len) {
        return MN_EARG;
    }
    
    size_t task;
    char *guard;
    
    struct timeval bt;
    gettimeofday(&bt, NULL);
    
    guard = buf;
    task = *len;
    while (task > 0) {
        size_t rst = recv(fd->sfd, guard, task, flags);
        if (0 == timeout) {
            if (rst > 0) {
                *len = rst;
                return 0;
            } else {
                *len = 0;
                return MN_ERECV;
            }
        }
        
        struct timeval st;
        gettimeofday(&st, NULL);
        if (timeval_min_usec(&st, &bt) > timeout) {
            if (rst >0) {
                *len = (*len) - (task - rst);
            } else {
                *len = (*len) - task;
            }
            return MN_ETIMEOUT;
        }
        if (rst > 0) {
            guard += rst;
            task -= rst;
        } else if (task == 0) {
            //close
            *len = (*len) - task;
            return 0;
        } else {
            if (errno==EINTR || errno==EAGAIN) {
                continue;
            } else {
                // other errors
                *len = (*len) - task;
                return MN_ESEND;
            }
        }
    }
    return 0;
}

