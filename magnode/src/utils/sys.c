//
//  sys.c
//  magnode
//
//  Created by apollo on 15/12/30.
//  Copyright © 2015年 proj-m. All rights reserved.
//

#include "sys.h"


#ifdef MN_ANDROID
uint64_t htonll(uint64_t val) {
    return (((uint64_t) htonl(val)) << 32) + htonl(val >> 32);
}

uint64_t ntohll(uint64_t val) {
    return (((uint64_t) ntohl(val)) << 32) + ntohl(val >> 32);
}
#endif