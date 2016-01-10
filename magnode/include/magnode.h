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
    
    typedef struct mn_node_t mn_node;
    
    typedef enum mn_key_type_t {
        MN_KEY_NONE,
        MN_KEY_SALT,
        MN_KEY_DH,
    } mn_key_type;
    
    typedef enum mn_protobuf_type_t {
        MN_PB_BIN,
        MN_PB_PB,
    } mn_protobuf_type;
    
    typedef enum mn_crypto_type_t {
        MN_CRYPTO_NONE,
        MN_CRYPTO_AES128,
    } mn_crypto_type;
    
    
    /**
     * New a Node.
     *
     * @return : a mn_node point
     */
    mn_node *mn_new();
    
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
     * Set channel infomation
     *
     * @param node: mn_node object to set channel info
     * @param protobuf : proto buffer type
     * @param key : key chain type
     * @param crypto : crypto method
     * @return : 0 on success , <0 on error
     */
    int mn_set_channel(mn_node *node, mn_protobuf_type protobuf, mn_key_type key, mn_crypto_type crypto);
    
    /**
     * Set Auth infomation
     *
     */
    int mn_set_auth(mn_node *node, const char *openid, const char *accesskey);
     
    /**
     * Connect to Server.
     *
     * @param node: a mn_node object
     * @param url: url to connect
     * @param timeout: connect timeout
     * @return : 0 on success , <0 on error
     */
    int mn_connect(mn_node *node,const char *url, uint32_t timeout);

    /**
     * Reconnect to Server.
     *
     * @param node: mn_node object to reconnect
     * @param timeout: timeout to reconnect
     * @return : 0 on success , <0 on error
     */
    int mn_reconnect(mn_node *node, uint32_t timeout);
    
    /**
     * Send Message Data.
     *
     * @param node: mn_node object for send
     * @param buf: message data buffer
     * @param length: message data buffer's length
     * @param timeout : send timout
     * @return : 0 on success , <0 on error
     */
    int mn_send(mn_node *node,const void *buf,size_t length,uint32_t timeout);
    
    /**
     * Recv Message Data.
     *
     * @param node: mn_node object for recv
     * @param buf: destination message data buffer
     * @param length: message data buffer's length
     * @param timeout : recv timout
     * @return : 0 on success , <0 on error
     */
    int mn_recv(mn_node *node,void *buf,size_t *length,uint32_t timeout);
    
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
