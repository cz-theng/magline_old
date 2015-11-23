//
//  main.m
//  test_magnode
//
//  Created by apollo on 15/9/23.
//  Copyright © 2015年 cz. All rights reserved.
//


#import <Foundation/Foundation.h>

#include "magnode.h"

void _main()
{
    mn_node node;
    mn_init(&node);
    //int ret = mn_connect(&node, "tcp://127.0.0.1:8082", 500*1000);
    int ret = mn_connect(&node, "tcp://123.57.145.218:8082", 500*1000);
    if (ret != 0){
        printf("connect error with ret %d",(int)ret);
    }
    
    printf("Connect Success!");
    
    size_t len = 5;
    int s = mn_send(&node, "nimei", len ,0);
    if (s<0) {
        printf("send error with size %d",(int)s);
    } else {
        printf("send with length %d",(int)len);
    }
    
    len = 5;
    char buf[1024] = {0};
    s = mn_recv(&node, buf, len, 0);
    if (s < 0) {
        printf("recv error with size %d", (int) len);
    }
    mn_close(&node);

}

int main(int argc, const char * argv[]) {
    
    _main();
    
    
    @autoreleasepool {
        // insert code here...
        NSLog(@"Hello, World!");
    }
    return 0;
}
