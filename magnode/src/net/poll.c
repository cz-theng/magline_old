/**
 * Author :cz cz.theng@gmail.com
 * Licence MIT
 */

#include "poll.h"
#include "net.h"


int mn_poll(int fd, int type, uint64_t timeout)
{
    if (0 == timeout) {
        return 0;
    }
    // should check less than FD_SETSIZE when use select
    fd_set rfds,wfds;
    int rst;
    struct timeval poll_timeout;
    poll_timeout.tv_sec = timeout/1000;
    poll_timeout.tv_usec= (timeout % 1000) *1000 ;
    
    FD_ZERO(&rfds);
    FD_SET(fd, &rfds);
    FD_ZERO(&wfds);
    FD_SET(fd, &wfds);
    
    if (MN_POLL_OUT == type) {
        rst = select(fd+1, NULL, &wfds, NULL, &poll_timeout);
    } else if (MN_POLL_IN == type) {
        rst = select(fd+1, &rfds, NULL, NULL, &poll_timeout);
    } else if (MN_POLL_INOUT == type) {
        rst = select(fd+1, &rfds, &wfds, NULL, &poll_timeout);
    }
    
    if( rst<0 ) {
        return MN__EPOLL;
    }
    
    if (0 == rst) {
        return MN__ETIMEOUT;
    }
    
    if (MN_POLL_IN==type && FD_ISSET(fd, &rfds)) {
        return 0;
    } else if (MN_POLL_OUT==type && FD_ISSET(fd, &wfds)) {
        return 0;
    } else if (MN_POLL_INOUT == type && FD_ISSET(fd, &rfds) &&FD_ISSET(fd, &wfds)) {
        return  0;
    }
    
    return MN__EPOLL;
}


