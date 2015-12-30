//
//  syn_ack.h
//  magnode
//
//  Created by apollo on 15/12/29.
//  Copyright © 2015年 proj-m. All rights reserved.
//

#ifndef syn_ack_h
#define syn_ack_h

#include "frame_head.h"
#pragma pack(1)
typedef  struct mn_ack_t
{
    mn_frame_head frame_head;
} mn_ack;
#pragma pack()



#pragma pack(1)
typedef struct mn_syn_t
{
    mn_frame_head frame_head;
} mn_syn;
#pragma pack()
#endif /* syn_ack_h */
