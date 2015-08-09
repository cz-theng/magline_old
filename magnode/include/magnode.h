//
//  magnode.h
//  magnode
//
//  Created by CZ on 7/23/15.
//  Copyright (c) 2015 proj-m. All rights reserved.
//

#ifndef magnode_magnode_h
#define magnode_magnode_h

#include <stdint.h>

#include "magnode_errcode.h"

#ifdef __cplusplus
extern "C" {
#endif
    
    typedef struct node_t {

    } mn_node;

    int mn_connect(mn_node *n,const char *url, uint64_t timeout);
    
    int mn_reconnect(mn_node *n, uint64_t timeout);
    
    int mn_send(mn_node *node,const void *buf,size_t length,uint64_t timeout);
    
    int mn_recv(mn_node *node,void *buf,size_t *length,uint64_t timeout);
    
    int mn_close(mn_node *node);

#ifdef __cplusplus
}
#endif

#endif