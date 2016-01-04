//
//  frame_head.c
//  magnode
//
//  Created by apollo on 15/12/29.
//  Copyright © 2015年 proj-m. All rights reserved.
//

// we all use little endian

#include "frame_head.h"
#include "proto.h"

int mn_init_frame_head(mn_frame_head *head, uint16_t cmd, uint32_t length)
{
    if (NULL == head) {
        return MN_EARG;
    }
    head->cmd = cmd;
    head->length = length;
    
    head->magic = MN_MAGIC;
    head->version = MN_VERSION;
    head->seq = tick_seq();
    return 0;
}

uint32_t tick_seq()
{
    static uint32_t seq = 0;
    return seq++;
}


int mn_pack_frame_head(mn_frame_head *head, void *buf, int len)
{
    if (NULL == head || NULL == buf) {
        return MN_EARG;
    }
    
    if (len < sizeof(*head)) {
        return MN_EPACKLEN;
    }
    
    memcpy(buf, head, sizeof(*head));
    
    return sizeof(*head);
}

int mn_unpack_frame_head(mn_frame_head *head, const void *buf, int len)
{
    if (NULL == head || NULL == buf) {
        return MN_EARG;
    }
    
    if (len < sizeof(*head)) {
        return MN_EPACKLEN;
    }
    
    memcpy(head, buf, sizeof(*head));
    return 0;
}