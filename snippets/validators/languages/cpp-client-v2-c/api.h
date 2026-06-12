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

static inline void LDUserSetFirstName(struct LDUser *user, const char *firstName) {
    (void)user;
    (void)firstName;
}

static inline void LDUserSetLastName(struct LDUser *user, const char *lastName) {
    (void)user;
    (void)lastName;
}

static inline void LDUserSetCustomAttributesJSON(struct LDUser *user,
                                                 struct LDJSON *custom) {
    (void)user;
    (void)custom;
}

/* LDJSON construction surface (shared c-json API) used by the
 * custom-attributes doc fragments. */
static inline struct LDJSON *LDNewObject(void) {
    return (struct LDJSON *)0;
}

static inline struct LDJSON *LDNewArray(void) {
    return (struct LDJSON *)0;
}

static inline struct LDJSON *LDNewText(const char *text) {
    (void)text;
    return (struct LDJSON *)0;
}

static inline LDBoolean LDArrayPush(struct LDJSON *array, struct LDJSON *item) {
    (void)array;
    (void)item;
    return LDBooleanTrue;
}

static inline LDBoolean LDObjectSetKey(struct LDJSON *object,
                                       const char *key,
                                       struct LDJSON *item) {
    (void)object;
    (void)key;
    (void)item;
    return LDBooleanTrue;
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

#ifdef __cplusplus
}
#endif

#endif /* LAUNCHDARKLY_API_H */
