//
//  magnode_inner.h
//  magnode
//
//  Created by apollo on 15/12/4.
//  Copyright © 2015年 proj-m. All rights reserved.
//

#ifndef magnode_inner_h
#define magnode_inner_h


typedef struct node_t {
    struct mn_socket socket;
    uint32_t agent_id;
    void *sendbuf;
    size_t sendbuflen;
    void *recvbuf;
    size_t recvbuflen;
} mn_node;

#endif /* magnode_inner_h */
