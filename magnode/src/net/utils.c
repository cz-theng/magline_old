/**
 * Author :cz cz.theng@gmail.com
 * Licence MIT
 */
#include "os.h"
#if defined MN_APPLE  || defined MN_ANDROID
#include <arpa/inet.h>
#endif

#include "utils.h"






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


