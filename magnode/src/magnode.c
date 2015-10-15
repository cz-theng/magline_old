//
//  magnode.c
//  magnode
//
//  Created by apollo on 15/9/23.
//  Copyright © 2015年 proj-m. All rights reserved.
//

#include "magnode.h"
#include "magnode_errcode.h"


int mn_init(mn_node *node)
{
    return -1;
}

int mn_deinit(mn_node *node)
{
    return -1;
}

int mn_connect(mn_node *node,const char *url, uint64_t timeout)
{
    return -1;
}

int mn_reconnect(mn_node *node, uint64_t timeout)
{
    return -1;
}

int mn_send(mn_node *node,const void *buf,size_t length,uint64_t timeout)
{
    return -1;
}

int mn_recv(mn_node *node,void *buf,size_t *length,uint64_t timeout)
{
    return -1;
}

int mn_close(mn_node *node)
{
    return -1;
}