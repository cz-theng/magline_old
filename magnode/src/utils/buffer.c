//
//  buffer.c
//  magnode
//
//  Created by apollo on 15/12/30.
//  Copyright © 2015年 proj-m. All rights reserved.
//

#include "buffer.h"
#include "magnode_errcode.h"

#include <stddef.h>
#include <stdlib.h>
#include <string.h>

int mn_buffer_init(mn_buffer *buffer, int size)
{
    if (NULL == buffer) {
        return MN_EARG;
    }
    
    buffer->data = (void *) malloc(size);
    if (NULL == buffer->data) {
        mn_buffer_deinit(buffer);
        return MN_ENEWBUF;
    }
    memset(buffer->data, 0, size);
    buffer->length = size;
    buffer->used = 0;
    
    return 0;
}

int mn_buffer_reset(mn_buffer *buffer, int size)
{
    if (NULL == buffer) {
        return MN_EARG;
    }
    if (size < buffer->length) {
        memset(buffer->data, 0, buffer->length);
        buffer->length = size;
        buffer->used = 0;
        return 0;
    }
    
    buffer->data = (void *) realloc(buffer->data, size);
    if (NULL == buffer->data) {
        mn_buffer_deinit(buffer);
        return MN_ENEWBUF;
    }
    memset(buffer->data, 0, size);
    buffer->length = size;
    buffer->used = 0;
    
    return 0;
}

int mn_buffer_deinit(mn_buffer *buffer)
{
    if (NULL == buffer) {
        return MN_EARG;
    }
    if (buffer->data) {
        free(buffer->data);
    }
    
    buffer->length  = 0;
    buffer->used    = 0;
    
    return 0;
}

int mn_buffer_align(mn_buffer *buffer, int index)
{
    if (NULL == buffer) {
        return MN_EARG;
    }
    if (0 == index) {
        return 0;
    }
    memmove(buffer->data, buffer->data+index, buffer->used-index);
    buffer->used = buffer->used - index;
    return 0;
}