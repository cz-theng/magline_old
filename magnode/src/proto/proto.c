//
//  proto.c
//  test_magnode
//
//  Created by apollo on 15/10/16.
//  Copyright © 2015年 cz. All rights reserved.
//

#include <stdlib.h>
#include <string.h>

#include "proto.h"
#include "magnode_errcode.h"



int parse2mem(nodemsg_head *head, const void *body, size_t body_len, void *buf, size_t buflen)
{
    size_t idx = 0;
    if (NULL == head || NULL == body || NULL == buf) {
        return MN_EARG;
    }
    if (buflen < (sizeof(nodemsg_head)+body_len)) {
        return  MN_EBODYLEN;
    }

    *(uint8_t *)((char *)buf+idx) = head->magic;
    idx++;
    *(uint8_t *)((char *)buf+idx) = head->version;
    idx++;
    *(uint16_t *)((char *)buf+idx) = htons(head->cmd);
    idx += 2;
    *(uint64_t *)((char *)buf+idx) = htonll(head->seq);
    idx += 4;
    *(uint64_t *)((char *)buf+idx) = htonll(head->agent_id);
    idx += 4;
    *(uint64_t *)((char *)buf+idx) = htonll(head->length);
    idx += 4;

    memcpy((void *)((char *)buf+idx), body, body_len);

    return 0;
}

int parse_from_mem(nodemsg_head *head, const void *body, void *buf, size_t *buflen)
{
    size_t idx = 0;
    if (NULL == head || NULL == body || NULL == buf) {
        return MN_EARG;
    }
    head->magic =*(uint8_t *)((char *)buf+idx);
    idx++;
    head->version =*(uint8_t *)((char *)buf+idx);
    idx++;
    head->cmd =ntohs(*(uint16_t *)((char *)buf+idx));
    idx += 2;
    head->seq = ntohll(*(uint64_t *)((char *)buf+idx));
    idx += 4;
    head->agent_id = ntohll(*(uint64_t *)((char *)buf+idx));
    idx += 4;
    head->length = ntohll(*(uint64_t *)((char *)buf+idx));
    idx += 4;
    
    if (head->length > *buflen) {
        return MN_EBODYLEN;
    }
    memcpy(buf, body, head->length);
    
    return 0;
}