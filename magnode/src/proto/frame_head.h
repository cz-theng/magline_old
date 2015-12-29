//
//  frame_head.h
//  magnode
//
//  Created by apollo on 15/12/29.
//  Copyright © 2015年 proj-m. All rights reserved.
//

#ifndef frame_head_h
#define frame_head_h

#include <stdint.h>

#pragma pack(1)

typedef struct mn_frame_head_t {
    uint8_t magic;
    uint8_t version;
    uint16_t cmd;
    uint32_t seq;
    uint32_t agent_id;
    uint16_t length;

    
} mn_frame_head;

#pragma pack()

#endif /* frame_head_h */
