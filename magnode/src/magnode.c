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

#if defined MN_APPLE  || defined MN_ANDROID
#include <sys/time.h>
#endif
#include <stdlib.h>
#include <string.h>



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
    rst = parse2mem(&head, NULL, 0, buf, buflen);
    if (0 != rst) {
        return -1;
    }
    struct timeval sbtime;
    gettimeofday(&sbtime, NULL);
    rst = mn_net_send(&node->socket, &head, &headlen, timeout);
    if (rst != 0 ) {
        return -1;
    }
    struct timeval setime;
    gettimeofday(&setime, NULL);
    long diff = timeval_min_usec(&setime, &sbtime);
    if (diff <= 0) {
        return MN_ETIMEOUT;
    }
    uint64_t rtimeout = diff;
    
    buflen = sizeof(mn_nodemsg_head);
    memset(buf, 0, buflen);
    gettimeofday(&sbtime, NULL);
    rst = mn_net_recv(&node->socket, buf, &buflen, rtimeout);
    if (rst != 0 ) {
        if (rst == MN__ETIMEOUT) {
            return MN_ETIMEOUT;
        }
        return -1;
    }
    diff = timeval_min_usec(&setime, &sbtime);
    if (diff <= 0) {
        return MN_ETIMEOUT;
    }
    rst = parse_from_mem(&head, NULL, &buflen, buf);
    if (rst != 0 ){
        return -1;
    }
    
    rst = is_invalied_head(&head);
    if (rst != 0) {
        return MN_EHEAD;
    }
    if (head.cmd == MN_CMD_RSP_CONN) {
        return 0;        
    } else {
        return MN_ECMD;
    }
    
    return -1;
}

int mn_init(mn_node *node)
{
    if (NULL == node){
        return MN_EARG;
    }
    return 0;
}

int mn_deinit(mn_node *node)
{
    if (NULL == node){
        return MN_EARG;
    }
    return 0;
}

int mn_connect(mn_node *node,const char *url, uint64_t timeout)
{
    int rst;
    
    if (NULL == node || NULL==url ){
        return MN_EARG;
    }
    struct timeval cbtime;
    gettimeofday(&cbtime, NULL);
    rst = mn_net_connect(url, &node->socket, timeout);
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
    if (diff <= 0) {
        return MN_ETIMEOUT;
    }
    uint64_t rtimeout = diff;
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
        return MN_EARG;
    }
    return 0;
}

int mn_send(mn_node *node,const void *buf,size_t length,uint64_t timeout)
{
    if (NULL == node || NULL == buf ){
        return MN_EARG;
    }
    return 0;
}

int mn_recv(mn_node *node,void *buf,size_t length,uint64_t timeout)
{
    if (NULL == node || NULL == buf){
        return MN_EARG;
    }
    return 0;
}

int mn_close(mn_node *node)
{
    if (NULL == node){
        return MN_EARG;
    }
    return 0;
}