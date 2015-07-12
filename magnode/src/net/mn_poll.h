//
//  mn_poll.h
//  magnode
//
//  Created by cz on 15/7/9.
//  Copyright (c) 2015年cz. All rights reserved.
//

#ifndef __magnode__mn_poll__
#define __magnode__mn_poll__

#include <sys/select.h>
#include <stdint.h>
#include <stddef.h>
#include <errno.h>


    
enum PollType {
    MN_POLL_IN  = 1,
    MN_POLL_OUT = 2,
};
    
int mn_poll(int fd, int type, uint64_t timeout);


#endif /* defined(__magnode__mn_poll__) */
