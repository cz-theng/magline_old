//
//  magnode_inner.c
//  magnode
//
//  Created by apollo on 15/12/29.
//  Copyright © 2015年 proj-m. All rights reserved.
//

#include "magnode_inner.h"
#include "syn_ack.h"
#include "magnode.h"
#include "proto.h"


int mn_clear_legacy_sendbuf(mn_node *node)
{
    if (NULL == node) {
        return MN_ENULLNODE;
    }
    
    if ( 0 ==  node->sendbuf.length ) {
        return 0;
    }
    size_t ts = node->sendbuf.length;
    int rst = mn_net_send(&node->socket, node->sendbuf.data, &ts, 0);
    if (ts == node->sendbuf.length) {
        node->sendbuf.length = 0;
    } else {
        mn_buffer_align(&node->sendbuf,(int) ts);
    }
    if (rst) {
        return rst;
    }

    return 0;
}

int mn_unpack_legacy_recvbuf(mn_node *node)
{
    if (NULL == node) {
        return MN_ENULLNODE;
    }
    
    if ( 0 ==  node->recvbuf.length ) {
        return 0;
    }
    
    return 0;
}

int mn_send_packbuf(mn_node *node)
{
    int rst = 0;
    if (NULL == node) {
        return MN_ENULLNODE;
    }
    
    if ((node->sendbuf.cap - node->sendbuf.length)<node->packbuf.length) {
        return MN_EPACKLEN;
    }
    rst = mn_buffer_append(&node->sendbuf, &node->packbuf);
    if (rst <0) {
        return rst;
    }
    rst = mn_clear_legacy_sendbuf(node);
    if (rst < 0) {
        return rst;
    }
    return 0;
}

uint32_t mn_cal_remain_time(struct timeval begintime, uint32_t timeout)
{
    uint32_t remain = 0;
    struct timeval now;
    gettimeofday(&now, NULL);
    long elapse = timeval_min_usec(&now, &begintime);
    remain = timeout - (uint32_t)elapse;
    remain = remain>0 ? remain : 0;
    return remain;
}

int mn_send_syn(mn_node *node, uint32_t timeout)
{
    int rst = 0;
    if (NULL == node) {
        return MN_ENULLNODE;
    }
    
    rst = mn_clear_legacy_sendbuf(node);
    if (rst) {
        LOG_E("Clear send buffer's legacy error");
        return rst;
    }
    
    mn_syn syn;
    rst = mn_init_syn(&syn, MN_PB_BIN, MN_KEY_NONE, MN_CRYPTO_NONE);
    if (rst <0) {
        LOG_E("init syn error with %d", rst);
        return rst;
    }
    
    mn_buffer_reset(&node->packbuf, MN_MAX_PROTO_SIZE);
    mn_pack_syn(&syn, node->packbuf.data, node->packbuf.length);
    rst = mn_send_packbuf(node);
    if (rst < 0) {
        LOG_E("send syn error with %d", rst);
        return rst;
    }
    return 0;
}



int mn_recv_ack(mn_node *node, uint32_t timeout)
{
    int rst =0;
    
    if (NULL == node) {
        return MN_ENULLNODE;
    }
    
    rst = mn_unpack_legacy_recvbuf(node);
    if (rst < 0) {
        return rst;
    }
    
    return 0;
}

int mn_send_session_req(mn_node *node, uint32_t timeout)
{
    if (NULL == node) {
        return MN_ENULLNODE;
    }
    
    return 0;
}

int mn_recv_session_rsp(mn_node *node, uint32_t timeout)
{
    if (NULL == node) {
        return MN_ENULLNODE;
    }
    
    return 0;
}

int mn_send_auth_req(mn_node *node, uint32_t timeout)
{
    if (NULL == node) {
        return MN_ENULLNODE;
    }
    return 0;
}

int mn_recv_auth_rsp(mn_node *node, uint32_t timeout)
{
    if (NULL == node) {
        return MN_ENULLNODE;
    }
    return 0;
}

int mn_recv_confirm(mn_node *node, uint32_t timeout)
{
    if (NULL == node) {
        return MN_ENULLNODE;
    }
    return 0;
}

int mn_connect_transaction(mn_node *node, uint32_t timeout)
{
    int rst = 0;
    uint32_t rt = timeout;
    
    if (NULL == node) {
        return MN_ENULLNODE;
    }
    
    struct timeval btime;
    gettimeofday(&btime, NULL);
    // 1. send syn
    rst = mn_send_syn(node, rt);
    if (rst <0 ) {
        LOG_E("send syn error");
        if (MN_ETIMEOUT == rst ) {
            return rst;
        } else {
            return MN_ESYN;
        }
    }
    
    // 2. recv ack
    rt = mn_cal_remain_time(btime, timeout);
    rst = mn_recv_ack(node, rt);
    if (rst <0 ) {
        LOG_E("send syn error");
        if (MN_ETIMEOUT == rst ) {
            return rst;
        } else {
            return MN_EACK;
        }
    }
    
    // send connect req
    
    // recv connect rsp
    
    // send auth
    
    // recv auth
    
    // recv confirm
    
#if 0
    MN_NODEMSG_HEAD_INIT(&head, MN_CMD_REQ_CONN, 0);
    headlen = sizeof(mn_nodemsg_head);
    
    size_t buflen = sizeof(mn_nodemsg_head);
    void *buf = malloc(sizeof(mn_nodemsg_head));
    rst = parse2mem(&head, NULL, 0, buf, &buflen);
    if (0 != rst) {
        return -1;
    }
    struct timeval sbtime;
    gettimeofday(&sbtime, NULL);
    rst = mn_net_send(&node->socket, buf, &headlen, timeout);
    if (rst != 0 ) {
        FREE(buf);
        return -1;
    }
    struct timeval setime;
    gettimeofday(&setime, NULL);
    long diff = timeval_min_usec(&setime, &sbtime);
    if (diff < 0 || diff > timeout) {
        FREE(buf);
        return MN_ETIMEOUT;
    }
    uint64_t rtimeout = timeout-diff;
    
    buflen = sizeof(mn_nodemsg_head);
    memset(buf, 0, buflen);
    gettimeofday(&sbtime, NULL);
    rst = mn_net_recv(&node->socket, buf, &buflen, rtimeout);
    if (rst != 0 ) {
        if (rst == MN__ETIMEOUT) {
            FREE(buf);
            return MN_ETIMEOUT;
        }
        FREE(buf);
        return -1;
    }
    gettimeofday(&setime, NULL);
    diff = timeval_min_usec(&setime, &sbtime);
    if (diff < 0 || diff > rtimeout) {
        FREE(buf);
        return MN_ETIMEOUT;
    }
    buflen = 0;
    rst = parse_from_mem(&head, NULL, &buflen, buf);
    if (rst != 0 ){
        FREE(buf);
        return -1;
    }
    
    rst = is_invalied_head(&head);
    if (rst != 0) {
        FREE(buf);
        return MN_EHEAD;
    }
    if (head.cmd == MN_CMD_RSP_CONN) {
        FREE(buf);
        node->agent_id =head.agent_id;
        return 0;
    } else {
        FREE(buf);
        return MN_ECMD;
    }
#endif
    return -1;
}

