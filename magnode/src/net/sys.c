//
//  sys.c
//  magnode
//
//  Created by apollo on 15/12/15.
//  Copyright © 2015年 proj-m. All rights reserved.
//

#include "sys.h"


#if defined MN_APPLE  || defined MN_ANDROID
#include <signal.h>
#endif

void mn_sys_ignore_pipe(void)
{
#if defined(MN_APPLE) || defined(MN_ANDROID)
    return;
#else
    struct sigaction act;
    
    sigemptyset(&act.sa_mask);
    act.sa_handler = SIG_IGN;
    act.sa_flags = 0;
    
    sigaction(SIGPIPE, &act, NULL);
#endif
}
