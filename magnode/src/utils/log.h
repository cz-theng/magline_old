//
//  cdnv_log.h
//  cdnvister
//
//  Created by cz on 15/7/7.
//  Copyright (c) 2015年cz. All rights reserved.
//

#ifndef MAGNODE_UTILS_MN_LOG_H_
#define MAGNODE_UTILS_MN_LOG_H_


#ifdef __cplusplus
extern "C"
{
#endif
    
    

    enum logger_level
    {
        LDEBUG=1,
        LINFO,
        LWARNING,
        LERROR,
        LFATAL,
    };
    
    void mn_log_init();
    
    void mn_log(int level, const char *file,int line,const char *func,const char *fmt, ...);
    
    void mn_log_set_level(enum logger_level level);
    
    
    #define  LOG_INIT() mn_log_init();
    #define LOG_D(...) mn_log(LDEBUG, __FILE__, __LINE__, __FUNCTION__, __VA_ARGS__);
    #define LOG_I(...) mn_log(LINFO, __FILE__, __LINE__, __FUNCTION__, __VA_ARGS__);
    #define LOG_W(...) mn_log(LWARNING, __FILE__, __LINE__, __FUNCTION__, __VA_ARGS__);
    #define LOG_E(...) mn_log(LERROR, __FILE__, __LINE__, __FUNCTION__, __VA_ARGS__);
    #define LOG_F(...) mn_log(LFATAL, __FILE__, __LINE__, __FUNCTION__, __VA_ARGS__);
        
    #define LOG_SET_LEVEL(level) mn_log_set_level((level));
    
#ifdef __cplusplus
}
#endif
    
#endif /* defined(MAGNODE_UTILS_MN_LOG_H_) */
