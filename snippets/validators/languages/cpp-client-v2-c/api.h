/*
 * Stub of <launchdarkly/api.h> from the v2.x C client SDK
 * (launchdarkly/c-client-sdk — distinct from c-server-sdk). Declares
 * the v2 client API surface that doc fragments under
 * sdk-docs/*-c-sdk-v2-x-native*.snippet.md reference.
 *
 * Definitions are inline no-ops — this header is for parse +
 * type-check validation only, never linked into a running binary.
 */
#ifndef LAUNCHDARKLY_API_H
#define LAUNCHDARKLY_API_H

#include <stdbool.h>
#include <stddef.h>

#ifdef __cplusplus
extern "C" {
#endif

typedef bool LDBoolean;
#define LDBooleanTrue  ((LDBoolean)1)
#define LDBooleanFalse ((LDBoolean)0)

struct LDConfig;
struct LDClient;
struct LDUser;
struct LDJSON;

static inline struct LDConfig *LDConfigNew(const char *key) {
    (void)key;
    return (struct LDConfig *)0;
}

static inline struct LDUser *LDUserNew(const char *key) {
    (void)key;
    return (struct LDUser *)0;
}

static inline void LDUserFree(struct LDUser *user) {
    (void)user;
}

/* cpp-client v2 C SDK's `LDClientInit` takes (config, user, maxwait) —
 * the user is bound at init time (mobile/client SDK pattern), unlike
 * the server SDK where the user is passed per-variation call. */
static inline struct LDClient *LDClientInit(struct LDConfig *config,
                                            struct LDUser *user,
                                            unsigned int maxwait) {
    (void)config;
    (void)user;
    (void)maxwait;
    return (struct LDClient *)0;
}

static inline LDBoolean LDClientClose(struct LDClient *client) {
    (void)client;
    return LDBooleanTrue;
}

/* Variation calls take (client, flagKey, fallback) — no per-call
 * user since the client carries the user. Distinct from the server
 * variant which takes (client, user, flagKey, fallback, &details). */
static inline LDBoolean LDBoolVariation(struct LDClient *client,
                                        const char *flagKey,
                                        LDBoolean fallback) {
    (void)client;
    (void)flagKey;
    return fallback;
}

static inline int LDIntVariation(struct LDClient *client,
                                 const char *flagKey,
                                 int fallback) {
    (void)client;
    (void)flagKey;
    return fallback;
}

static inline double LDDoubleVariation(struct LDClient *client,
                                       const char *flagKey,
                                       double fallback) {
    (void)client;
    (void)flagKey;
    return fallback;
}

static inline char *LDStringVariation(struct LDClient *client,
                                      const char *flagKey,
                                      const char *fallback) {
    (void)client;
    (void)flagKey;
    return (char *)fallback;
}

#ifdef __cplusplus
}
#endif

#endif /* LAUNCHDARKLY_API_H */
