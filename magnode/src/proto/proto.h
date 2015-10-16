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

const uint8_t MAGIC = 0x7f;
const uint8_t VERSION = 0x01;

enum CMD {
    CMD_REQ_CONN = 0x0001,
    CMD_RSP_CONN = 0x0002,
    
    CMD_REQ_SEND = 0x0003,
    CMD_RSP_SEND = 0x0004,

    CMD_REQ_RECV = 0x0005,
    CMD_RSP_RECV = 0x0006,
    
    CMD_REQ_CLOSE = 0x0007,
    CMD_RSP_CLOSE = 0x0008,
    
    CMD_REQ_RECONN = 0x0009,
    CMD_RSPREQCONN = 0x000a,

};

typedef struct nodemsg_head_t {
    uint8_t magic;
    uint8_t version;
    uint16_t cmd;
    uint64_t seq;
    uint64_t agent_id;
    uint64_t length;
} nodemsg_head;

int parse2mem(nodemsg_head *head, const void *body, size_t body_len, void *buf, size_t buflen);

int parse_from_mem(nodemsg_head *head, const void *body, void *buf, size_t *buflen);

uint64_t tick_seq()
{
    static uint64_t seq = 0;
    return seq++;
}

#endif /* proto_h */
