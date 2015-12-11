//
//  magnode_uinty.c
//  magnode
//
//  Created by apollo on 15/11/17.
//  Copyright © 2015年 proj-m. All rights reserved.
//

#include <stdio.h>
#include <stdlib.h>
#include "log.h"

#include "magnode_uinty.h"

void *exp_create()
{
    LOG_I("exp_create");
    mn_node *node = mn_new();
    return node;
}

int exp_mn_init(void *node)
{
    LOG_I("exp_mn_init");
    int rst;
    if (!node) {
        return -1;
    }
    mn_node *n = (mn_node *)node;
    rst = mn_init(n);
    return rst;
}

int exp_mn_deinit(void *node)
{
    LOG_I("exp_mn_deinit");
    int rst;
    if (!node) {
        return -1;
    }
    mn_node *n = (mn_node *)node;
    rst = mn_deinit(n);
    return rst;
}

int exp_mn_connect(void *node, char *url, int timeout)
{
    LOG_I("exp_mn_connect");
    int rst;
    if (!node) {
        return -1;
    }
    mn_node *n = ( mn_node *)node;
    rst = mn_connect(n, url, timeout);
    return rst;
}

int exp_mn_send(void *node, char *data, int length, int timeout)
{
    LOG_I("exp_mn_send");
    int rst;
    if (!node) {
        return -1;
    }
    mn_node *n = (mn_node *)node;
    rst = mn_send(n, data, length, timeout);
    return rst;
}

int exp_mn_recv(void *node, char *data, int *length, int timeout)
{
    LOG_I("exp_mn_recv");
    int rst;
    size_t len = *length;
    if (!node) {
        return -1;
    }
    mn_node *n = (mn_node *)node;
    rst = mn_recv(n, data, &len, timeout);
    if (rst != 0) {
        return -1;
    }
    *length = len;
    LOG_I("length is %d ",*length);
    return rst;
}

int exp_mn_close(void *node)
{
    LOG_I("exp_mn_close");
    int rst;
    if (!node) {
        return -1;
    }
    mn_node *n = (mn_node *)node;
    rst = mn_close(n);
    return rst;
}

void exp_mn_destory(void *node)
{
    LOG_I("exp_mn_destory");
    if (!node) {
        return ;
    }
    mn_node *n = (mn_node *)node;
    mn_deinit(n);
    free(n);
}