//
//  sys.h
//  magnode
//
//  Created by apollo on 15/12/30.
//  Copyright © 2015年 proj-m. All rights reserved.
//

#ifndef sys_h
#define sys_h

#include "os.h"

#if defined MN_APPLE  || defined MN_ANDROID
#include <arpa/inet.h>
#endif


#ifdef MN_ANDROID
uint64_t htonll(uint64_t val);
uint64_t ntohll(uint64_t val);
#endif

//#define UINT16(buf) ()

#endif /* sys_h */
