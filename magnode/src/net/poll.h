/**
 * Author :cz cz.theng@gmail.com
 * Licence MIT
 */

#ifndef MAGNODE_NET_MN_POLL_H_
#define MAGNODE_NET_MN_POLL_H_

#include "os.h"

#if defined  MN_APPLE || defined MN_ANDROID
#include <sys/select.h>
#endif

#if defined  MN_WIN
#include <winsock2.h>
#endif
#include <stdint.h>
#include <stddef.h>
#include <errno.h>


    
enum poll_type {
    MN_POLL_IN  = 0x01,
    MN_POLL_OUT = 0x02,
    MN_POLL_INOUT = 0x04,
};

// type could be xor by poll_type
int mn_poll(int fd, int type, uint64_t timeout);


#endif /* defined(MAGNODE_NET_MN_POLL_H_) */
