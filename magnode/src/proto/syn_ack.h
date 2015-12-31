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
typedef  struct mn_syn_t
{
    mn_frame_head frame_head;
    uint16_t protobuf;
    uint16_t key;
    uint16_t crypto;
    uint32_t seq;
} mn_syn;
#pragma pack()


#pragma pack(1)
typedef struct mn_ack_none_t
{
    
} mn_ack_none;
#pragma pack()

#pragma pack(1)
typedef struct mn_ack_salt_t
{
    
} mn_ack_salt;
#pragma pack()

#pragma pack(1)
typedef struct mn_ack_dh_t
{
    
} mn_ack_dh;
#pragma pack()

#pragma pack(1)
typedef union mn_ack_body_t
{
    mn_ack_none ack_none;
    mn_ack_salt ack_salt;
    mn_ack_dh   ack_dh;
} mn_ack_body;
#pragma pack()

#pragma pack(1)
typedef struct mn_axk_t
{
    mn_frame_head frame_head;
    uint16_t key;
    mn_ack_body body;
} mn_ack;
#pragma pack()


#ifdef __cplusplus
extern "C" {
#endif
    int mn_init_syn(mn_syn *syn, uint16_t protobuf, uint16_t key, uint16_t crypto);
    
    int mn_pack_syn(mn_syn *syn, void *buf, int len);
    
    int mn_unpack_syn(mn_syn *syn, const void *buf, int len);
    
    int mn_pack_ack(mn_ack *ack, void *buf, int len);
    
    int mn_unpack_ack(mn_ack *ack, const void *buf, int len);
    
#ifdef __cplusplus
}
#endif
#endif /* syn_ack_h */
