//
//  buffer.h
//  magnode
//
//  Created by apollo on 15/12/30.
//  Copyright © 2015年 proj-m. All rights reserved.
//

#ifndef buffer_h
#define buffer_h

typedef  struct mn_buffer_t {
    void *data;
    int length;
    int cap;
} mn_buffer;

#ifdef __cplusplus
extern "C" {
#endif

    int mn_buffer_init(mn_buffer *buffer, int size);

    int mn_buffer_reset(mn_buffer *buffer, int size);

    int mn_buffer_deinit(mn_buffer *buffer);

    int mn_buffer_align(mn_buffer *buffer, int index);
    
    int mn_buffer_append(mn_buffer *dest, mn_buffer *src);
    
#ifdef __cplusplus
}
#endif

#endif /* buffer_h */
