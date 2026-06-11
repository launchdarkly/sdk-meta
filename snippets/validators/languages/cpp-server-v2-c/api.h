/*
 * Stub of <launchdarkly/api.h> from the v2.x C server SDK
 * (launchdarkly/c-server-sdk). Declares the v2 API surface that doc
 * fragments under sdk-docs/c-c---*-c-sdk-v2-x*.snippet.md reference.
 * Definitions are inline no-ops — this header is for parse + type-check
 * validation only, never linked into a running binary.
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

/* Evaluation-reason surface. Mirrors the real v2 header's enum order
 * and the LDDetails fields the doc fragments touch (the real struct
 * carries an additional per-kind `extra` union the fragments never
 * reference). */
enum LDEvalReason {
    LD_UNKNOWN = 0,
    LD_ERROR,
    LD_OFF,
    LD_PREREQUISITE_FAILED,
    LD_TARGET_MATCH,
    LD_RULE_MATCH,
    LD_FALLTHROUGH
};

struct LDDetails {
    unsigned int variationIndex;
    LDBoolean hasVariation;
    enum LDEvalReason reason;
};

static inline void LDDetailsInit(struct LDDetails *details) {
    (void)details;
}

static inline void LDDetailsClear(struct LDDetails *details) {
    (void)details;
}

static inline struct LDConfig *LDConfigNew(const char *key) {
    (void)key;
    return (struct LDConfig *)0;
}

static inline struct LDClient *LDClientInit(struct LDConfig *config, unsigned int maxwait) {
    (void)config;
    (void)maxwait;
    return (struct LDClient *)0;
}

static inline LDBoolean LDClientClose(struct LDClient *client) {
    (void)client;
    return LDBooleanTrue;
}

static inline struct LDUser *LDUserNew(const char *key) {
    (void)key;
    return (struct LDUser *)0;
}

static inline void LDUserFree(struct LDUser *user) {
    (void)user;
}

static inline void LDUserSetName(struct LDUser *user, const char *name) {
    (void)user;
    (void)name;
}

static inline void LDUserSetEmail(struct LDUser *user, const char *email) {
    (void)user;
    (void)email;
}

static inline void LDUserSetFirstName(struct LDUser *user, const char *firstName) {
    (void)user;
    (void)firstName;
}

static inline void LDUserSetLastName(struct LDUser *user, const char *lastName) {
    (void)user;
    (void)lastName;
}

static inline void LDUserSetAnonymous(struct LDUser *user, LDBoolean anonymous) {
    (void)user;
    (void)anonymous;
}

static inline LDBoolean LDBoolVariation(struct LDClient *client,
                                        struct LDUser *user,
                                        const char *flagKey,
                                        LDBoolean fallback,
                                        struct LDDetails *details) {
    (void)client;
    (void)user;
    (void)flagKey;
    (void)details;
    return fallback;
}

static inline int LDIntVariation(struct LDClient *client,
                                 struct LDUser *user,
                                 const char *flagKey,
                                 int fallback,
                                 struct LDDetails *details) {
    (void)client;
    (void)user;
    (void)flagKey;
    (void)details;
    return fallback;
}

static inline double LDDoubleVariation(struct LDClient *client,
                                       struct LDUser *user,
                                       const char *flagKey,
                                       double fallback,
                                       struct LDDetails *details) {
    (void)client;
    (void)user;
    (void)flagKey;
    (void)details;
    return fallback;
}

static inline char *LDStringVariation(struct LDClient *client,
                                      struct LDUser *user,
                                      const char *flagKey,
                                      const char *fallback,
                                      struct LDDetails *details) {
    (void)client;
    (void)user;
    (void)flagKey;
    (void)details;
    return (char *)fallback;
}

#ifdef __cplusplus
}
#endif

#endif /* LAUNCHDARKLY_API_H */
