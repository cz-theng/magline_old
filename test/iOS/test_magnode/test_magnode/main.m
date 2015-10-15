//
//  main.m
//  test_magnode
//
//  Created by apollo on 15/9/23.
//  Copyright © 2015年 cz. All rights reserved.
//


#import <Foundation/Foundation.h>

#include "mn_net.h"
#include "mn_log.h"

void _main()
{
    struct mn_socket socket_hdl;
    
    int a = strlen(NULL);
    
    int ret = mn_connect("tcp://127.0.0.1:8088", &socket_hdl, 5000);
    if (ret != 0){
        LOG_E("connect error with ret %d",(int)ret);
    }
    
    LOG_I("Connect Success!");
    
    size_t len = 5;
    int s = mn_send(&socket_hdl, "nimei", &len ,0);
    if (s<0) {
        LOG_E("send error with size %d",(int)s);
    } else {
        LOG_I("send with length %d",(int)len);
    }
    
    len = 5;
    char buf[1024] = {0};
    s = mn_recv(&socket_hdl, buf, &len, 0);
    if (s < 0) {
        LOG_E("recv error with size %d", (int) len);
    }
    mn_close(&socket_hdl);

}

int main(int argc, const char * argv[]) {
    
    _main();
    
    
    @autoreleasepool {
        // insert code here...
        NSLog(@"Hello, World!");
    }
    return 0;
}
