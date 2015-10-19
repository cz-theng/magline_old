//
//  magnode.h
//  magnode
//
//  Created by CZ on 7/23/15.
//  Copyright (c) 2015 proj-m. All rights reserved.
//

#ifndef magnode_magnode_h
#define magnode_magnode_h

#include <stdint.h>
#if defined __APPLE__ || defined __ANDROID__
#include <sys/socket.h>

#endif

#if defined  _WIN32 || _WIN64
#include <winsock2.h>
#include <WS2tcpip.h>
#endif


#ifdef __cplusplus
extern "C" {
#endif
    
    struct mn_socket
    {
        int sfd;
        int proto;
        struct sockaddr dest_addr;
        socklen_t addrlen;
    };
    
    typedef struct node_t {
        struct mn_socket socket;
        uint32_t agent_id;
        void *sendbuf;
        size_t sendbuflen;
        void *recvbuf;
        size_t recvbuflen;
    } mn_node;

    /**
     * Init a Node.
     *
     * @param node: mn_node object to init
     * @return : 0 on success , <0 on error
     */
    int mn_init(mn_node *node);
    
    /**
     * Deinit a Node.
     *
     * @param node: mn_node object to deinit
     * @return : 0 on success , <0 on error
     */
    int mn_deinit(mn_node *node);
    
     
    /**
     * Connect to Server.
     *
     * @param node: a mn_node object
     * @param url: url to connect
     * @param timeout: connect timeout
     * @return : 0 on success , <0 on error
     */
    int mn_connect(mn_node *node,const char *url, uint64_t timeout);

    /**
     * Reconnect to Server.
     *
     * @param node: mn_node object to reconnect
     * @param timeout: timeout to reconnect
     * @return : 0 on success , <0 on error
     */
    int mn_reconnect(mn_node *node, uint64_t timeout);
    
    /**
     * Send Message Data.
     *
     * @param node: mn_node object for send
     * @param buf: message data buffer
     * @param length: message data buffer's length
     * @param timeout : send timout
     * @return : 0 on success , <0 on error
     */
    int mn_send(mn_node *node,const void *buf,size_t length,uint64_t timeout);
    
    /**
     * Recv Message Data.
     *
     * @param node: mn_node object for recv
     * @param buf: destination message data buffer
     * @param length: message data buffer's length
     * @param timeout : recv timout
     * @return : 0 on success , <0 on error
     */
    int mn_recv(mn_node *node,void *buf,size_t length,uint64_t timeout);
    
    /**
     * Close Connection.
     *
     * @param node: mn_node to close
     * @return : 0 on success , <0 on error
     */
    int mn_close(mn_node *node);

#ifdef __cplusplus
}
#endif

#endif
