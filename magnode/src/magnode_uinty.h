//
//  magnode_uinty.h
//  magnode
//
//  Created by apollo on 15/11/17.
//  Copyright © 2015年 proj-m. All rights reserved.
//

#ifndef magnode_uinty_h
#define magnode_uinty_h

#include "magnode.h"

#if defined(WIN32) || defined(_WIN32)
#define MAGNODE_API __declspec(dllexport)
#else
#define MAGNODE_API
#endif

#ifdef __cplusplus
extern "C" {
#endif
    
    MAGNODE_API int exp_create();

    MAGNODE_API int exp_mn_init();
    
    MAGNODE_API int exp_mn_deinit();

    MAGNODE_API int exp_mn_connect(char *url, int timeout);
    
    MAGNODE_API int exp_mn_send(char *data, int length);
    
    MAGNODE_API int exp_mn_recv(char *data, int length);
    
    MAGNODE_API int exp_mn_close();
    
    
#ifdef __cplusplus
}
#endif

#endif /* magnode_uinty_h */
