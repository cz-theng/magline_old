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

// we all use little endian

enum {
    MN_MAX_MSG_LEN = 1024,
    MN_MAX_SENDBUF_SIZE = 10*1024,
    MN_MAX_RECVBUF_SIZE = 10*1024,
    MN_MAX_PROTO_SIZE = 10*1024,
    MN_MAX_TIMEOUT = 1000*1000,
    
    MN_MAGIC = 0x7f,
    MN_VERSION = 0x01,
};

enum MN_CMD {
    MN_CMD_UNKNOWN  = 0x0000,
    
    MN_CMD_SYN      = 0x0001,
    MN_CMD_ACK      = 0x0002,

};

typedef struct mn_nodemsg_head_t {
    uint8_t magic;
    uint8_t version;
    uint16_t cmd;
    uint32_t seq;
    uint32_t agent_id;
    uint32_t length;
} mn_nodemsg_head;

#undef MN_NODEMSG_HEAD_INIT
#define MN_NODEMSG_HEAD_INIT(head, CMD, id) do { (head)->magic = MN_MAGIC; \
                                        (head)->version = MN_VERSION; \
                                        (head)->cmd = (CMD); \
                                        (head)->seq = tick_seq(); \
                                        (head)->agent_id = (id); \
                                        (head)->length = 0; }while(0)

int parse2mem(mn_nodemsg_head *head, const void *body, size_t body_len, void *buf, size_t *buflen);

int parse_from_mem(mn_nodemsg_head *head, const void *body,size_t *bodylen, void *buf);

int is_invalied_head(mn_nodemsg_head *head);
#endif /* proto_h */
