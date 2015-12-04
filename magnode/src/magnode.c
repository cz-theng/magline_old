//
//  magnode.c
//  magnode
//
//  Created by apollo on 15/9/23.
//  Copyright © 2015年 proj-m. All rights reserved.
//

#include "os.h"
#include "magnode.h"
#include "magnode_errcode.h"
#include "net.h"
#include "utils.h"
#include "proto.h"
#include "log.h"
#include "magnode_inner.h"

#if defined MN_APPLE  || defined MN_ANDROID
#include <sys/time.h>
#endif
#include <stdlib.h>
#include <string.h>

#define FREE(p) do { if (NULL != p){free(p); p=NULL;} } while(0)

int connect_transaction(mn_node *node, uint64_t timeout)
{
    mn_nodemsg_head head;
    size_t headlen = 0;
    int rst = 0;

    
    if (NULL == node) {
        return -1;
    }
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
    
    return -1;
}

mn_node *mn_new()
{
    void *buf =(void *) malloc(sizeof(mn_node));
    memset(buf,0,sizeof(mn_node));
    mn_node *node = (mn_node *)buf;
    return node;
}

int mn_init(mn_node *node)
{
    if (NULL == node){
        LOG_E("mn_init: node is NULL");
        return MN_EARG;
    }
    LOG_I("mn_init(node %p)", node);
    
    node->agent_id = 0;
    node->sendbuf = malloc(MN_MAX_SENDBUF_SIZE);
    if (NULL == node->sendbuf) {
        LOG_E("Alloc Memory Error: malloc(MN_MAX_SENDBUF_SIZE)");
        return MN_EALLOC;
    }
    node->sendbuflen = MN_MAX_SENDBUF_SIZE;
    
    node->recvbuf = malloc(MN_MAX_RECVBUF_SIZE);
    if (NULL == node->recvbuf) {
        LOG_E("Alloc Memory Error: malloc(MN_MAX_RECVBUF_SIZE)");
        return MN_EALLOC;
    }
    node->recvbuflen = MN_MAX_RECVBUF_SIZE;
    return 0;
}

int mn_deinit(mn_node *node)
{
    if (NULL == node){
        LOG_E("mn_deinit: node is NULL");
        return MN_EARG;
    }
    LOG_I("mn_deinit(node %p)", node);
    
    if (NULL != node->sendbuf) {
        FREE(node->sendbuf);
        node->sendbuflen = 0;
    }
    if (NULL != node->recvbuf) {
        FREE(node->recvbuf);
        node->recvbuflen = 0;
    }
    node->agent_id = 0;
    return 0;
}

int mn_connect(mn_node *node,const char *url, uint64_t timeout)
{
    int rst;
    if (NULL == node || NULL==url ){
        LOG_E("mn_connect: node is NULL or url is NULL");
        return MN_EARG;
    }
    LOG_I("mn_connect(node %p,url %s, timeout %llu)", node, url, timeout);
    
    struct timeval cbtime;
    gettimeofday(&cbtime, NULL);
    rst = mn_net_connect(url, &node->socket, timeout);
    LOG_I("net connect with %d rst", rst);
    if (rst != 0 ) {
        if (rst == MN__ETIMEOUT) {
            return MN_ETIMEOUT;
        } else {
            return MN_ECONN;
        }
        
    }
    struct timeval cetime;
    gettimeofday(&cetime, NULL);
    long diff = timeval_min_usec(&cetime, &cbtime);
    if (diff < 0 || diff > timeout) {
        LOG_I("mn_connect timeout with diff %ld",diff);
        return MN_ETIMEOUT;
    }
    uint64_t rtimeout = timeout - diff;
    rst = connect_transaction(node, rtimeout);
    if (rst < 0) {
        if (MN_ETIMEOUT ==rst ) {
            return MN_ETIMEOUT;
        } else {
            return MN_ECONN;
        }
    }
    return 0;
}

int mn_reconnect(mn_node *node, uint64_t timeout)
{
    if (NULL == node){
        LOG_E("mn_reconnect: node is NULL");
        return MN_EARG;
    }
    LOG_I("mn_reconnect(node %p, timeout %llu)", node, timeout);
    
    return 0;
}

int mn_send(mn_node *node,const void *buf,size_t length,uint64_t timeout)
{
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
    
    return 0;
}

int mn_recv(mn_node *node,void *buf,size_t length,uint64_t timeout)
{
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
        return node->recvbuflen;
    } else {
        return MN_ECMD;
    }
}

int mn_close(mn_node *node)
{
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
    
    return 0;
}