//
//  mn_log.cpp
//  magnode
//
//  Created by cz on 15/7/7.
//  Copyright (c) 2015年cz. All rights reserved.
//

#include "mn_log.h"

#include <stdarg.h>
#ifdef __APPLE__
#import <Foundation/Foundation.h>
#endif

#ifdef __ANDROID__
# include <android/log.h>
# define _TAG     	"ApolloCDNVister"
#endif


void mn_log(const char *fmt, ...)
{
    va_list args;
    va_start(args, fmt);


#ifdef __ANDROID__
    __android_log_print(ANDROID_LOG_DEBUG, _TAG, fmt, args);
#endif

#ifdef __APPLE__
    NSString *fmt_nsstr = [NSString stringWithCString:fmt
                                             encoding:[NSString defaultCStringEncoding]];
    NSLog(fmt_nsstr,args);
#endif
    va_end(args);
}


