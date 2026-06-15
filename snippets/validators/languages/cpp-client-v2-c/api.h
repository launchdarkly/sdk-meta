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
#include <stdio.h>
#include <string.h>

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

static inline void LDConfigSetUseEvaluationReasons(struct LDConfig *config,
                                                   LDBoolean reasons) {
    (void)config;
    (void)reasons;
}

/* Detail-variation surface. `LDVariationDetails` is filled by the
 * evaluation call; `reason` is a JSON object inspected through
 * LDObjectLookup / LDGetText. LDFreeDetailContents takes the struct
 * by value, matching the real v2 header. */
typedef struct {
    int            variationIndex;
    struct LDJSON *reason;
} LDVariationDetails;

static inline LDBoolean LDBoolVariationDetail(struct LDClient *client,
                                              const char *flagKey,
                                              LDBoolean fallback,
                                              LDVariationDetails *details) {
    (void)client;
    (void)flagKey;
    (void)details;
    return fallback;
}

static inline int LDIntVariationDetail(struct LDClient *client,
                                       const char *flagKey,
                                       int fallback,
                                       LDVariationDetails *details) {
    (void)client;
    (void)flagKey;
    (void)details;
    return fallback;
}

static inline double LDDoubleVariationDetail(struct LDClient *client,
                                             const char *flagKey,
                                             double fallback,
                                             LDVariationDetails *details) {
    (void)client;
    (void)flagKey;
    (void)details;
    return fallback;
}

static inline struct LDJSON *LDObjectLookup(const struct LDJSON *object,
                                            const char *key) {
    (void)object;
    (void)key;
    return (struct LDJSON *)0;
}

static inline const char *LDGetText(const struct LDJSON *json) {
    (void)json;
    return "";
}

static inline void LDFreeDetailContents(LDVariationDetails details) {
    (void)details;
}

/* Custom-event surface. The real v2 client header's LDClientTrack
 * takes just the client and the event key (the client already holds
 * the user); data/metric variants exist separately. */
static inline void LDClientTrack(struct LDClient *client, const char *key) {
    (void)client;
    (void)key;
}

/* All-flags surface. The real v2 header returns an object-type LDJSON
 * map of flag keys to values; iteration goes through the shared LDJSON
 * collection helpers below. */
static inline struct LDJSON *LDAllFlags(struct LDClient *client) {
    (void)client;
    return (struct LDJSON *)0;
}

static inline struct LDJSON *LDGetIter(const struct LDJSON *collection) {
    (void)collection;
    return (struct LDJSON *)0;
}

static inline struct LDJSON *LDIterNext(const struct LDJSON *iter) {
    (void)iter;
    return (struct LDJSON *)0;
}

static inline const char *LDIterKey(const struct LDJSON *iter) {
    (void)iter;
    return "";
}

static inline char *LDJSONSerialize(const struct LDJSON *json) {
    (void)json;
    return (char *)0;
}

static inline void LDFree(void *buffer) {
    (void)buffer;
}

static inline void LDJSONFree(struct LDJSON *json) {
    (void)json;
}

#ifdef __cplusplus
}
#endif

#endif /* LAUNCHDARKLY_API_H */
