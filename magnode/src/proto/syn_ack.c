//
//  syn_ack.c
//  magnode
//
//  Created by apollo on 15/12/29.
//  Copyright © 2015年 proj-m. All rights reserved.
//

#include "syn_ack.h"
#include "magnode.h"
#include "log.h"
#include "proto.h"

int mn_init_syn(mn_syn *syn, uint16_t protobuf, uint16_t key, uint16_t crypto)
{
    if (NULL == syn) {
        return MN_EARG;
    }
    syn->crypto = crypto;
    syn->key = key;
    syn->protobuf = protobuf;

    mn_init_frame_head(&syn->frame_head, MN_CMD_SYN, sizeof(*syn) - sizeof(syn->frame_head));
    return 0;
}

int mn_pack_syn(mn_syn *syn, void *buf, int len)
{
    int rst = 0;
    if (NULL == syn || NULL == buf) {
        return MN_EARG;
    }
    
    if (len < sizeof(*syn)) {
        return MN_EPACKLEN;
    }
    
    rst = mn_pack_frame_head(&syn->frame_head, buf, len);
    if (rst < 0) {
        return rst;
    }
    
    memcpy(buf+rst, syn, sizeof(*syn));
    return 0;
}

int mn_unpack_syn(mn_syn *syn, const void *buf, int len)
{
    return 0;
}

int mn_pack_ack(mn_ack *ack, void *buf, int len)
{
    return 0;
}

int mn_unpack_ack(mn_ack *ack, const void *buf, int len)
{
    if (NULL == ack || NULL == buf) {
        return MN_EARG;
    }
    
    if (len < sizeof(uint16_t)) {
        return MN_EPACKLEN;
    }
    
    uint16_t key = *((uint16_t *) buf);
    switch (key) {
        case MN_KEY_NONE: {
            memcpy(&ack->body.ack_none, buf+sizeof(key), sizeof(ack->body.ack_none));
        }
            break;
            
        case MN_KEY_DH: {
            memcpy(&ack->body.ack_dh, buf+sizeof(key), sizeof(ack->body.ack_dh));
        }
            break;
            
        case MN_KEY_SALT: {
            memcpy(&ack->body.ack_salt, buf+sizeof(key), sizeof(ack->body.ack_salt));
        }
            break;
            
        default: {
            LOG_E("Unapck ACK wiht unknown key");
            return MN_EUNKNOWNACK;
        }
            break;
    }
    
    return 0;
}