//
//  frame_head.c
//  magnode
//
//  Created by apollo on 15/12/29.
//  Copyright © 2015年 proj-m. All rights reserved.
//

#include "frame_head.h"

#include <stdlib.h>
#include <string.h>

// we all use little endian
int mn_pack_frame_head(mn_frame_head *head, void *buf, int len)
{
    if (NULL == head || buf == head) {
        return MN_EARG;
    }
    
    if (len < sizeof(*head)) {
        return MN_EPACKLEN;
    }
    
    memcpy(buf, head, sizeof(*head));
    
    return 0;
}

int mn_unpack_frame_head(mn_frame_head *head, const void *buf, int len)
{
    if (NULL == head || buf == head) {
        return MN_EARG;
    }
    
    if (len < sizeof(*head)) {
        return MN_EPACKLEN;
    }
    
    memcpy(head, buf, sizeof(*head));
    return 0;
}