//
//  magnode.c
//  magnode
//
//  Created by apollo on 15/9/23.
//  Copyright © 2015年 proj-m. All rights reserved.
//


#include "magnode_inner.h"

#if defined MN_APPLE  || defined MN_ANDROID
#include <sys/time.h>
#endif
#include <stdlib.h>
#include <string.h>



mn_node *mn_new()
{
    // just like golang's new . new with 0 memory
    void *buf =(void *) malloc(sizeof(mn_node));
    memset(buf,0,sizeof(mn_node));
    mn_node *node = (mn_node *)buf;
    return node;
}

int mn_init(mn_node *node)
{
    int rst = 0;
    if (NULL == node){
        LOG_E("mn_init: node is NULL");
        return MN_EARG;
    }
    LOG_I("mn_init(node %p)", node);
    
    node->agent_id = 0;
    
    rst = mn_buffer_init(&node->sendbuf, MN_MAX_PROTO_SIZE);
    if (rst ) {
        LOG_E("Init send buffer error");
        mn_deinit(node);
        return rst;
    }
    
    rst = mn_buffer_init(&node->recvbuf, MN_MAX_PROTO_SIZE *2 );
    if (rst ) {
        LOG_E("Init recv buffer error");
        mn_deinit(node);
        return rst;
    }
    
    rst = mn_buffer_init(&node->packbuf, MN_MAX_PROTO_SIZE );
    if (rst ) {
        LOG_E("Init pack buffer error");
        mn_deinit(node);
        return rst;
    }
    
    memset(&node->socket, 0, sizeof(node->socket));
    return 0;
}

int mn_deinit(mn_node *node)
{
    if (NULL == node){
        LOG_E("mn_deinit: node is NULL");
        return MN_EARG;
    }
    LOG_I("mn_deinit(node %p)", node);
    
    mn_buffer_deinit(&node->sendbuf);
    mn_buffer_deinit(&node->recvbuf);
    mn_buffer_deinit(&node->packbuf);
    node->agent_id = 0;
    return 0;
}

int mn_connect(mn_node *node,const char *url, uint32_t timeout)
{
    int rst;
    if (NULL == node || NULL==url ){
        LOG_E("mn_connect: node is NULL or url is NULL");
        return MN_EARG;
    }
    LOG_I("mn_connect(node %p,url %s, timeout %llu)", node, url, timeout);
    
    struct timeval btime;
    gettimeofday(&btime, NULL);
    rst = mn_net_connect(&node->socket, url, timeout);
    LOG_D("net connect with %d rst", rst);
    if (rst != 0 ) {
        if (rst == MN__ETIMEOUT) {
            return MN_ETIMEOUT;
        } else {
            return MN_ECONN;
        }
        
    }
    
    uint32_t rt = mn_cal_remain_time(btime, timeout);
    if ( 0 == rt) {
        LOG_I("mn_connect timeout");
        return MN_ETIMEOUT;
    }
    rst = mn_connect_transaction(node, rt);
    if (rst < 0) {
        if (MN_ETIMEOUT ==rst ) {
            return MN_ETIMEOUT;
        } else {
            return rst;
        }
    }
    
    return 0;
}

int mn_reconnect(mn_node *node, uint32_t timeout)
{
    if (NULL == node){
        LOG_E("mn_reconnect: node is NULL");
        return MN_EARG;
    }
    LOG_I("mn_reconnect(node %p, timeout %llu)", node, timeout);
    
    return 0;
}

int mn_send(mn_node *node,const void *buf,size_t length,uint32_t timeout)
{
#if 0
    mn_nodemsg_head head;
    int rst;
    if (NULL == node || NULL == buf ){
        LOG_E("mn_send: node is NULL or buf is NULL");
        return MN_EARG;
    }
    
    MN_NODEMSG_HEAD_INIT(&head, MN_CMD_MSG_NODE, node->agent_id);
    if (length < MN_MAX_SENDBUF_SIZE) {
        node->sendbuflen  = length;
    }
    node->sendbuflen =MN_MAX_SENDBUF_SIZE;
    head.agent_id = node->agent_id;
    head.length = length;
    rst = parse2mem(&head, buf, length, node->sendbuf, &node->sendbuflen);
    if (rst != 0) {
        return MN_EPARSE;
    }
    
    struct timeval sbtime;
    gettimeofday(&sbtime, NULL);
    rst = mn_net_send(&node->socket, node->sendbuf, &node->sendbuflen, timeout);
    if (rst != 0 ) {
        return MN_ESEND;
    }
    struct timeval setime;
    gettimeofday(&setime, NULL);
    long diff = timeval_min_usec(&setime, &sbtime);
    if (diff<0 || (timeout >0 &&diff > timeout)) {
        return MN_ETIMEOUT;
    }
#endif
    return 0;
}

int mn_recv(mn_node *node,void *buf,size_t *length,uint32_t timeout)
{
#if 0
    mn_nodemsg_head head;
    int rst;
    struct timeval sbtime;
    struct timeval setime;
    long diff;
    if (NULL == node || NULL == buf){
        LOG_E("mn_recv: node is NULL or buf is NULL");
        return MN_EARG;
    }
    
    MN_NODEMSG_HEAD_INIT(&head, MN_CMD_REQ_SEND, node->agent_id);
    
    node->recvbuflen = sizeof(mn_nodemsg_head);
    memset(node->recvbuf, 0, MN_MAX_RECVBUF_SIZE);
    gettimeofday(&sbtime, NULL);
    rst = mn_net_recv(&node->socket, node->recvbuf, &node->recvbuflen, timeout);
    if (rst != 0 ) {
        if (rst == MN__ETIMEOUT) {
            return MN_ETIMEOUT;
        }
        return MN_ERECV;
    }
    gettimeofday(&setime, NULL);
    diff = timeval_min_usec(&setime, &sbtime);
    if (diff < 0 ||(timeout >0 && diff > timeout)) {
        return MN_ETIMEOUT;
    }
    rst = parse_from_mem(&head, node->recvbuf, &node->recvbuflen, node->recvbuf);
    if (rst != 0 ){
        return MN_EUNPARSE;
    }
    
    rst = is_invalied_head(&head);
    if (0 != rst ) {
        return MN_EHEAD;
    }
    if (MN_CMD_MSG_KNOT) {
        uint64_t rtimeout = timeout - diff;
        memset(node->recvbuf, 0, MN_MAX_RECVBUF_SIZE);
        node->recvbuflen = head.length;
        gettimeofday(&sbtime, NULL);
        rst = mn_net_recv(&node->socket, node->recvbuf, &node->recvbuflen, rtimeout);
        if (rst != 0 ) {
            if (rst == MN__ETIMEOUT) {
                return MN_ETIMEOUT;
            }
            return MN_ERECV;
        }
        gettimeofday(&setime, NULL);
        if (diff < 0 || (timeout > 0 && diff > rtimeout)) {
            return MN_ETIMEOUT;
        }
        memcpy(buf, node->recvbuf, node->recvbuflen);
        *length =node->recvbuflen;
        return 0;
    } else {
        return MN_ECMD;
    }
#endif 
    return 0;
}

int mn_close(mn_node *node)
{
#if 0
    mn_nodemsg_head head;
    int rst;
    if (NULL == node){
        LOG_E("mn_close: node is NULL");
        return MN_EARG;
    }
    LOG_I("mn_close(node %p)", node);
    
    MN_NODEMSG_HEAD_INIT(&head, MN_CMD_REQ_CLOSE, node->agent_id);
    
    rst = parse2mem(&head, NULL, 0, node->sendbuf, &node->sendbuflen);
    if (rst != 0) {
        return MN_EPARSE;
    }
    
    rst = mn_net_send(&node->socket, node->sendbuf, &node->sendbuflen, MN_MAX_TIMEOUT);
    if (rst != 0 ) {
        return MN_ESEND;
    }
    mn_net_close(&node->socket);
#endif
    return 0;
}