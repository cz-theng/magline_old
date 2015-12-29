//
//  magnode_inner.c
//  magnode
//
//  Created by apollo on 15/12/29.
//  Copyright © 2015年 proj-m. All rights reserved.
//

#include "magnode_inner.h"


uint32_t cal_remain_time(struct timeval begintime, uint32_t timeout)
{
    uint32_t remain = 0;
    struct timeval now;
    gettimeofday(&now, NULL);
    long elapse = timeval_min_usec(&now, &begintime);
    remain = timeout - (uint32_t)elapse;
    remain = remain>0 ? remain : 0;
    return remain;
}

int connect_transaction(mn_node *node, uint32_t timeout)
{
    mn_nodemsg_head head;
    size_t headlen = 0;
    int rst = 0;
    
    
    if (NULL == node) {
        return -1;
    }
    
    
    // send syn
    
    // recv ack
    
    // send connect req
    
    // recv connect rsp
    
    // send auth
    
    // recv auth
    
    // recv confirm
    
    
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

