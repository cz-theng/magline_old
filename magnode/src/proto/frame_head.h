//
//  frame_head.h
//  magnode
//
//  Created by apollo on 15/12/29.
//  Copyright © 2015年 proj-m. All rights reserved.
//

#ifndef frame_head_h
#define frame_head_h

#include "magnode_errcode.h"
#include "sys.h"

#include <stdlib.h>
#include <string.h>
#include <stdint.h>

// we use little endian

#pragma pack(1)

typedef struct mn_frame_head_t
{
    uint8_t magic;
    uint8_t version;
    uint16_t cmd;
    uint32_t seq;
    uint32_t agent_id;
    uint32_t length;

    
} mn_frame_head;

#pragma pack()


#ifdef __cplusplus
extern "C" {
#endif
    int mn_init_frame_head(mn_frame_head *head, uint16_t cmd, uint32_t length);
    
    int mn_pack_frame_head(mn_frame_head *head, void *buf, int len);
    
    int mn_unpack_frame_head(mn_frame_head *head, const void *buf, int len);
    
    uint32_t tick_seq();
    
#ifdef __cplusplus
}
#endif

#endif /* frame_head_h */
