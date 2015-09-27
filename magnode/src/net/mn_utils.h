/**
 * Author :cz cz.theng@gmail.com
 * Licence MIT
 */

#ifndef MAGNODE_NET_MN_UTILS_H_
#define MAGNODE_NET_MN_UTILS_H_

#include "mn_os.h"

#if defined MN_APPLE  || defined MN_ANDROID
#include <sys/time.h>
#endif

#ifdef MN_ANDROID
uint64_t htonll(uint64_t val);
uint64_t ntohll(uint64_t val);
#endif

long timeval_cmp (const struct timeval *tv1, const struct timeval *tv2);
long timeval_min_usec(const struct timeval *tv1, const struct timeval *tv2);


#endif /* MAGNODE_NET_MN_UTILS_H_ */
