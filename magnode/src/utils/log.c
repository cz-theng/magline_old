//
//  cdnv_log.cpp
//  cdnvister
//
//  Created by cz on 15/7/7.
//  Copyright (c) 2015年cz. All rights reserved.
//

#include <stdio.h>
#include <pthread.h>

#include "log.h"
#include "os.h"


#include <stdarg.h>
#include <stdlib.h>
#include <string.h>
#include <stdio.h>

#ifdef MN_APPLE
//#include <Foundation/Foundation.h>
#endif

#ifdef MN_ANDROID
# include <android/log.h>
# define _TAG     	"Magnode"
#endif


enum {
    MN_LOG_BUF_SIZE = 1024*2,
};

typedef struct mn_logger_t
{
    enum logger_level level;
    char log_buf[MN_LOG_BUF_SIZE];
    pthread_mutex_t mutex;
} mn_logger;

static mn_logger g_logger;

#ifdef MN_MAC
FILE *g_fp = NULL;
#endif

void mn_print(const char *log)
{
#ifdef MN_ANDROID
    __android_log_print(ANDROID_LOG_DEBUG, _TAG, "%s", log);
#endif
    
#ifdef MN_IOS
    printf("%s\n",log);
#endif
#ifdef MN_MAC
    if (NULL == g_fp) {
        g_fp = fopen("/tmp/magline.log", "wb");
        if (NULL == g_fp) {
            printf("%s", log);
            return;
        }
    } else {
        fprintf(g_fp,"%s \n", log);
        fflush(g_fp);
    }
#endif
}

void mn_log_init()
{
    pthread_mutex_init(&g_logger.mutex, NULL);
}

void mn_log(int level, const char *file,int line,const char *func,const char *fmt, ...)
{
    if (level < g_logger.level) {
        return ;
    }
    va_list args;
    va_start(args, fmt);
    
    pthread_mutex_lock(&g_logger.mutex);
    memset(g_logger.log_buf, 0, MN_LOG_BUF_SIZE);
    int index = snprintf(g_logger.log_buf, MN_LOG_BUF_SIZE, "[%s(%d) %s()]:", file, line, func);
    vsnprintf(g_logger.log_buf+index, MN_LOG_BUF_SIZE-index, fmt,args);
    mn_print(g_logger.log_buf);
    pthread_mutex_lock(&g_logger.mutex);
    va_end(args);
}

void mn_log_set_level(enum logger_level level)
{
    g_logger.level = level;
}
