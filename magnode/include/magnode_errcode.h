//
//  magnode_errcode.h
//  magnode
//
//  Created by CZ on 7/23/15.
//  Copyright (c) 2015 proj-m. All rights reserved.
//

#ifndef magnode_magnode_errcode_h
#define magnode_magnode_errcode_h

#ifdef __cplusplus
extern  "C" {
#endif
    
#define MN_ERRNO_BASE 10000
    
    enum mn_errno {
        EARG = -(MN_ERRNO_BASE +1),
    };
    
#ifdef __cplusplus
}
#endif

#endif
