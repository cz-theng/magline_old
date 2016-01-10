//
//  proto.c
//  test_magnode
//
//  Created by apollo on 15/10/16.
//  Copyright © 2015年 cz. All rights reserved.
//

#include <stdlib.h>
#include <string.h>
#include "os.h"
#if defined MN_APPLE  || defined MN_ANDROID
#include <arpa/inet.h>
#endif
#include "proto.h"
#include "magnode_errcode.h"



int parse2mem(mn_nodemsg_head *head, const void *body, size_t body_len, void *buf, size_t *buflen)
{
    size_t idx = 0;
    if (NULL == head || NULL == buf || NULL == buflen) {
        return MN_EARG;
    }
    if (*buflen < (sizeof(mn_nodemsg_head)+body_len)) {
        return  MN_EBODYLEN;
    }

    *(uint8_t *)((char *)buf+idx) = head->magic;
    idx++;
    *(uint8_t *)((char *)buf+idx) = head->version;
    idx++;
    *(uint16_t *)((char *)buf+idx) = htons(head->cmd);
    idx += 2;
    *(uint32_t *)((char *)buf+idx) = htonl(head->seq);
    idx += 4;
    *(uint32_t *)((char *)buf+idx) = htonl(head->agent_id);
    idx += 4;
    *(uint32_t *)((char *)buf+idx) = htonl(head->length);
    idx += 4;

    memcpy((void *)((char *)buf+idx), body, body_len);
    *buflen = body_len + sizeof(mn_nodemsg_head);
    return 0;
}

int parse_from_mem(mn_nodemsg_head *head, const void *body,size_t *bodylen, void *buf)
{
    size_t idx = 0;
    if (NULL == head || NULL == buf) {
        return MN_EARG;
    }
    //MN_NODEMSG_HEAD_INIT(head, MN_CMD_UNKNOWN, 0);
    
    head->magic =*(uint8_t *)((char *)buf+idx);
    idx++;
    head->version =*(uint8_t *)((char *)buf+idx);
    idx++;
    head->cmd =ntohs(*(uint16_t *)((char *)buf+idx));
    idx += 2;
    head->seq = ntohl(*(uint32_t *)((char *)buf+idx));
    idx += 4;
    head->agent_id = ntohl(*(uint32_t *)((char *)buf+idx));
    idx += 4;
    head->length = ntohl(*(uint32_t *)((char *)buf+idx));
    idx += 4;
    
    if (head->length > *bodylen) {
        return MN_EBODYLEN;
    }
    memcpy(buf, body, head->length);
    *bodylen = head->length;
    return 0;
}



int is_invalied_head(mn_nodemsg_head *head) {
    if (NULL == head) {
        return -1;
    }
    
    if (head->magic == MN_MAGIC
        && head->version == MN_VERSION) {
    
        return 0;
    }
    return -1;
}