/**
 * Author :cz cz.theng@gmail.com
 * Licence MIT
 */

#include "mn_utils.h"

#ifdef MN_ANDROID
uint64_t htonll(uint64_t val) {
    return (((uint64_t) htonl(val)) << 32) + htonl(val >> 32);
}

uint64_t ntohll(uint64_t val) {
    return (((uint64_t) ntohl(val)) << 32) + ntohl(val >> 32);
}
#endif



long timeval_cmp (const struct timeval *tv1, const struct timeval *tv2)
{
    if (NULL == tv1 || NULL == tv2) {
        return -2;
    }
    return (tv1->tv_sec == tv2->tv_sec) ? (tv1->tv_usec - tv2->tv_usec) : (tv1->tv_sec - tv2->tv_sec);
}

long timeval_min_usec(const struct timeval *tv1, const struct timeval *tv2)
{
    if (NULL == tv1 || NULL == tv2) {
        return 0;
    }
    return (tv1->tv_sec*1000 + tv1->tv_usec/1000)-(tv2->tv_sec*1000 + tv2->tv_usec/1000);
}


