//
//  mn_poll.cpp
//  magnode
//
//  Created by cz on 15/7/9.
//  Copyright (c) 2015年cz. All rights reserved.
//

#include "mn_poll.h"
#include "mn_net.h"


int mn_poll(int fd, int type, uint64_t timeout)
{
    fd_set fds;
    int rst;
    struct timeval poll_timeout;
    poll_timeout.tv_sec = timeout/1000;
    poll_timeout.tv_usec= (timeout % 1000) *1000 ;
    
    FD_ZERO(&fds);
    FD_SET(fd, &fds);
    
    if (MN_POLL_OUT == type) {
        rst = select(fd+1, &fds, NULL, NULL, &poll_timeout);
    } else if (MN_POLL_IN == type) {
        rst = select(fd+1, NULL, &fds, NULL, &poll_timeout);
    }
    
    if( rst<0 ) {
        return MN_EPOLL;
    }
    
    if (0 == rst) {
        return MN_ETIMEOUT;
    }
    
    if (FD_ISSET(fd, &fds)) {
        return 0;
    }
    
    return MN_EPOLL;
}


