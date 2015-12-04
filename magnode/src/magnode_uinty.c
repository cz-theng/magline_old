//
//  magnode_uinty.c
//  magnode
//
//  Created by apollo on 15/11/17.
//  Copyright © 2015年 proj-m. All rights reserved.
//

#include <stdio.h>
#include <stdlib.h>

#include "magnode_uinty.h"

void *exp_create()
{
    printf("exp_create");
    mn_node *node = mn_new();
    return node;
}

int exp_mn_init(void *node)
{
    printf("exp_mn_init");
    if (!node) {
        return -1;
    }
    mn_node *n = (mn_node *)node;
    mn_init(n);
    return 0;
}

int exp_mn_deinit(void *node)
{
    printf("exp_mn_deinit");
    if (!node) {
        return -1;
    }
    mn_node *n = (mn_node *)node;
    mn_deinit(n);
    return 0;
}

int exp_mn_connect(void *node, char *url, int timeout)
{
    printf("exp_mn_connect");
    if (!node) {
        return -1;
    }
    mn_node *n = ( mn_node *)node;
    mn_connect(n, url, timeout);
    return 0;
}

int exp_mn_send(void *node, char *data, int length, int timeout)
{
    printf("exp_mn_send");
    if (!node) {
        return -1;
    }
    mn_node *n = (mn_node *)node;
    mn_send(n, data, length, timeout);
    return 0;
}

int exp_mn_recv(void *node, char *data, int length, int timeout)
{
    printf("exp_mn_recv");
    if (!node) {
        return -1;
    }
    mn_node *n = (mn_node *)node;
    mn_recv(n, data, length, timeout);
    return 0;
}

int exp_mn_close(void *node)
{
    printf("exp_mn_close");
    if (!node) {
        return -1;
    }
    mn_node *n = (mn_node *)node;
    mn_close(n);
    return 0;
}

void exp_mn_destory(void *node)
{
    if (!node) {
        return ;
    }
    mn_node *n = (mn_node *)node;
    mn_deinit(n);
    free(n);
}