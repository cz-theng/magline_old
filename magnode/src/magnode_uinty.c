//
//  magnode_uinty.c
//  magnode
//
//  Created by apollo on 15/11/17.
//  Copyright © 2015年 proj-m. All rights reserved.
//

#include <stdio.h>

#include "magnode_uinty.h"

static mn_node g_node ;

int exp_create()
{
    printf("exp_create");
    return 0;
}

int exp_mn_init()
{
    printf("exp_mn_init");
    mn_init(&g_node);
    return 0;
}

int exp_mn_deinit()
{
    printf("exp_mn_deinit");
    mn_deinit(&g_node);
    return 0;
}

int exp_mn_connect(char *url, int timeout)
{
    printf("exp_mn_connect");
    mn_connect(&g_node, url, timeout);
    return 0;
}

int exp_mn_send(char *data, int length)
{
    printf("exp_mn_send");
    mn_send(&g_node, data, length, 5000);
    return 0;
}

int exp_mn_recv(char *data, int length)
{
    printf("exp_mn_recv");
    mn_recv(&g_node, data, length, 5000);
    return 0;
}

int exp_mn_close()
{
    printf("exp_mn_close");
    mn_close(&g_node);
    return 0;
}