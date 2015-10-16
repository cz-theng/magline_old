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
        MN_EARG         = -(MN_ERRNO_BASE + 1),
        MN_ETIMEOUT     = -(MN_ERRNO_BASE + 2),
        MN_ECONN        = -(MN_ERRNO_BASE + 3),
        MN_EBODYLEN     = -(MN_ERRNO_BASE + 4),
        MN_EHEAD        = -(MN_ERRNO_BASE + 5),
        MN_ECMD         = -(MN_ERRNO_BASE + 6),
        MN_EALLOC       = -(MN_ERRNO_BASE + 7),
        MN_ESEND        = -(MN_ERRNO_BASE + 8),
        MN_ERECV        = -(MN_ERRNO_BASE + 9),
        MN_EPARSE       = -(MN_ERRNO_BASE + 10),
        MN_EUNPARSE     = -(MN_ERRNO_BASE + 11),
    };
    
#ifdef __cplusplus
}
#endif

#endif
