//
//  magnode_inner.h
//  magnode
//
//  Created by apollo on 15/12/4.
//  Copyright © 2015年 proj-m. All rights reserved.
//

#ifndef magnode_inner_h
#define magnode_inner_h

#include "os.h"
#include "magnode.h"
#include "magnode_errcode.h"
#include "net.h"
#include "utils.h"
#include "proto.h"
#include "log.h"

#if defined MN_APPLE  || defined MN_ANDROID
#include <sys/time.h>
#endif
#include <stdlib.h>
#include <string.h>

#include <stdint.h>
#if defined __APPLE__ || defined __ANDROID__
#include <sys/socket.h>

#endif

#if defined  _WIN32 || _WIN64
#include <winsock2.h>
#include <WS2tcpip.h>
#endif


#define FREE(p) do { if (NULL != p){free(p); p=NULL;} } while(0)

typedef struct node_t {
    struct mn_socket socket;
    uint32_t agent_id;
    void *sendbuf;
    size_t sendbuflen;
    void *recvbuf;
    size_t recvbuflen;
} mn_node;

#ifdef __cplusplus
extern "C" {
#endif
    
    uint32_t cal_remain_time(struct timeval begintime, uint32_t timeout);
    
    
    int connect_transaction(mn_node *node, uint32_t timeout);
    
#ifdef __cplusplus
}
#endif

#endif /* magnode_inner_h */
