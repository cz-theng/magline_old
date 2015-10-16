//
//  magnode.c
//  magnode
//
//  Created by apollo on 15/9/23.
//  Copyright © 2015年 proj-m. All rights reserved.
//

#include "os.h"
#include "magnode.h"
#include "magnode_errcode.h"
#include "net.h"
#include "utils.h"

#if defined MN_APPLE  || defined MN_ANDROID
#include <sys/time.h>
#endif
#include <stdlib.h>

int connect_transaction(uint64_t timeout)
{
    return -1;
}

int mn_init(mn_node *node)
{
    if (NULL == node){
        return MN_EARG;
    }
    return 0;
}

int mn_deinit(mn_node *node)
{
    if (NULL == node){
        return MN_EARG;
    }
    return 0;
}

int mn_connect(mn_node *node,const char *url, uint64_t timeout)
{
    int rst;
    
    if (NULL == node || NULL==url ){
        return MN_EARG;
    }
    struct timeval cbtime;
    gettimeofday(&cbtime, NULL);
    rst = mn_net_connect(url, &node->socket, timeout);
    if (rst != 0 ) {
        if (rst == MN__ETIMEOUT) {
            return MN_ETIMEOUT;
        } else {
            return MN_ECONN;
        }
        
    }
    struct timeval cetime;
    gettimeofday(&cetime, NULL);
    uint64_t rtimeout = labs(timeval_min_usec(&cetime, &cbtime));
    rst = connect_transaction(rtimeout);
    if (rst < 0) {
        return MN_ECONN;
    }
    return 0;
}

int mn_reconnect(mn_node *node, uint64_t timeout)
{
    if (NULL == node){
        return MN_EARG;
    }
    return 0;
}

int mn_send(mn_node *node,const void *buf,size_t length,uint64_t timeout)
{
    if (NULL == node || NULL == buf ){
        return MN_EARG;
    }
    return 0;
}

int mn_recv(mn_node *node,void *buf,size_t length,uint64_t timeout)
{
    if (NULL == node || NULL == buf){
        return MN_EARG;
    }
    return 0;
}

int mn_close(mn_node *node)
{
    if (NULL == node){
        return MN_EARG;
    }
    return 0;
}