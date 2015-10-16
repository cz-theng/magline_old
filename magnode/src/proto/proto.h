//
//  proto.h
//  test_magnode
//
//  Created by apollo on 15/10/16.
//  Copyright © 2015年 cz. All rights reserved.
//

#ifndef proto_h
#define proto_h

#include <stdio.h>
#include <stdint.h>

typedef struct nodemsg_head_t {
    uint8_t magic;
    uint8_t version;
    uint16_t cmd;
    uint64_t seq;
    uint64_t agent_id;
    uint64_t length;
} nodemsg_head;

int parse2mem(nodemsg_head *head, const void *body, size_t body_len, void *buf, size_t buflen);

int parse_from_mem(nodemsg_head *head, void *body, size_t *body_len, const void *buf, size_t buflen);

#endif /* proto_h */
