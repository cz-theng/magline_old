/**
 * Author :cz cz.theng@gmail.com
 * Licence MIT
 */

#ifndef MAGNODE_NET_MN_OS_H_
#define MAGNODE_NET_MN_OS_H_

/** Windows */
#if defined _WIN32
# define MN_WIN
# ifdef _WIN64
#  define MN_WIN64
# else 
#  define MN_WIN32
# endif
#endif

/** Apple */
#if defined __APPLE__
# define MN_APPLE
# include "TargetConditionals.h"
# if TARGET_IPHONE_SIMULATOR
#  define MN_IOS
#  define MN_IOS_SIMULATOR
# elif TARGET_OS_IPHONE
#  define MN_IOS
#  define MN_IOS_DEVICE
# elif TARGET_OS_MAC
#  define MN_MAC
# endif
#endif

/** Android */
#if defined __ANDROID__
# define MN_ANDROID
#endif


#endif /* defined(MAGNODE_NET_MN_OS_H_) */
