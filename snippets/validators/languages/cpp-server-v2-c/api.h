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

/* Logging surface. Mirrors the real v2 header's
 * <launchdarkly/logging.h> (vendored from c-sdk-common): the level
 * enum order, the convenience LDBasicLogger, the global logger
 * registration hook, and the level-to-string helper. */
typedef enum
{
    LD_LOG_FATAL = 0,
    LD_LOG_CRITICAL,
    LD_LOG_ERROR,
    LD_LOG_WARNING,
    LD_LOG_INFO,
    LD_LOG_DEBUG,
    LD_LOG_TRACE
} LDLogLevel;

static inline void LDBasicLogger(const LDLogLevel level, const char *const text) {
    (void)level;
    (void)text;
}

static inline void LDConfigureGlobalLogger(
    const LDLogLevel level,
    void (*logger)(const LDLogLevel level, const char *const text)) {
    (void)level;
    (void)logger;
}

static inline const char *LDLogLevelToString(const LDLogLevel level) {
    (void)level;
    return "";
}

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

/* Manual event flush: fire-and-forget; delivery happens on the
 * SDK's background thread. Matches the real v2 header's
 * `void LDClientFlush(struct LDClient *const client)`. */
static inline void LDClientFlush(struct LDClient *client) {
    (void)client;
}

static inline struct LDUser *LDUserNew(const char *key) {
    (void)key;
    return (struct LDUser *)0;
}

static inline void LDUserFree(struct LDUser *user) {
    (void)user;
}

/* Generates an identify event for the user, adding it to the Contexts
 * list without requiring a flag evaluation. */
static inline void LDClientIdentify(struct LDClient *client,
                                    const struct LDUser *user) {
    (void)client;
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

static inline void LDUserSetCustom(struct LDUser *user, struct LDJSON *custom) {
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

/* Custom-event surface. The real v2 header's LDClientTrack takes the
 * event key, the user the event is for, and optional JSON data. */
static inline void LDClientTrack(struct LDClient *client,
                                 const char *key,
                                 const struct LDUser *user,
                                 struct LDJSON *data) {
    (void)client;
    (void)key;
    (void)user;
    (void)data;
}

/* All-flags surface. The real v2 header returns an object-type LDJSON
 * map of flag keys to values; doc fragments only bind the result. */
static inline struct LDJSON *LDAllFlags(struct LDClient *client,
                                        const struct LDUser *user) {
    (void)client;
    (void)user;
    return (struct LDJSON *)0;
}

#ifdef __cplusplus
}
#endif

#endif /* LAUNCHDARKLY_API_H */
