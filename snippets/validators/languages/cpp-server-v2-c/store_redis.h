/*
 * Stub of <launchdarkly/store/redis.h> from the v2.x C server SDK
 * (launchdarkly/c-server-sdk). Declares the Redis store-integration
 * surface that the storing-data doc fragments reference. Definitions
 * are inline no-ops — this header is for parse + type-check validation
 * only, never linked into a running binary.
 */
#ifndef LAUNCHDARKLY_STORE_REDIS_H
#define LAUNCHDARKLY_STORE_REDIS_H

#include <launchdarkly/api.h>

#ifdef __cplusplus
extern "C" {
#endif

struct LDRedisConfig;

static inline struct LDRedisConfig *LDRedisConfigNew(void) {
    return (struct LDRedisConfig *)0;
}

static inline struct LDStoreInterface *LDStoreInterfaceRedisNew(
    struct LDRedisConfig *config) {
    (void)config;
    return (struct LDStoreInterface *)0;
}

#ifdef __cplusplus
}
#endif

#endif /* LAUNCHDARKLY_STORE_REDIS_H */
