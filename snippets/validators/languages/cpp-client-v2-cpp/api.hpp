/*
 * Stub of <launchdarkly/api.hpp> from the v2.x cpp-client SDK's
 * C++ binding (LDClientCPP). Declares the v2 C++ binding API surface
 * referenced by doc fragments under
 * sdk-docs/*-c-sdk-v2-x-c-binding*.snippet.md.
 *
 * Definitions are inline no-ops — header is for parse + type-check
 * validation only, never linked into a running binary.
 */
#ifndef LAUNCHDARKLY_API_HPP
#define LAUNCHDARKLY_API_HPP

#include <cstdbool>
#include <cstddef>
#include <cstring>
#include <iostream>

/* C-flavored types live in the global namespace, matching the
 * v2 SDK's actual layout (the C and C++ headers shared a translation
 * unit). */
typedef bool LDBoolean;
#define LDBooleanTrue  ((LDBoolean)1)
#define LDBooleanFalse ((LDBoolean)0)

struct LDConfig;
struct LDUser;
struct LDJSON;

inline LDConfig *LDConfigNew(const char *key) {
    (void)key;
    return nullptr;
}

inline LDUser *LDUserNew(const char *key) {
    (void)key;
    return nullptr;
}

inline void LDUserFree(LDUser *user) {
    (void)user;
}

inline void LDUserSetAnonymous(LDUser *user, LDBoolean anon) {
    (void)user;
    (void)anon;
}

inline void LDConfigSetUseEvaluationReasons(LDConfig *config,
                                            LDBoolean reasons) {
    (void)config;
    (void)reasons;
}

/* Detail-variation surface shared with the C API. `reason` is a JSON
 * object inspected through LDObjectLookup / LDGetText.
 * LDFreeDetailContents takes the struct by value, matching the real
 * v2 header. */
struct LDVariationDetails {
    int     variationIndex;
    LDJSON *reason;
};

inline LDJSON *LDObjectLookup(const LDJSON *object, const char *key) {
    (void)object;
    (void)key;
    return nullptr;
}

inline const char *LDGetText(const LDJSON *json) {
    (void)json;
    return "";
}

inline void LDFreeDetailContents(LDVariationDetails details) {
    (void)details;
}

/* All-flags surface. getAllFlags on LDClientCPP returns an object-type
 * LDJSON map of flag keys to values; iteration goes through the shared
 * LDJSON collection helpers. */
inline LDJSON *LDGetIter(const LDJSON *collection) {
    (void)collection;
    return nullptr;
}

inline LDJSON *LDIterNext(const LDJSON *iter) {
    (void)iter;
    return nullptr;
}

inline const char *LDIterKey(const LDJSON *iter) {
    (void)iter;
    return "";
}

inline char *LDJSONSerialize(const LDJSON *json) {
    (void)json;
    return nullptr;
}

inline void LDFree(void *buffer) {
    (void)buffer;
}

inline void LDJSONFree(LDJSON *json) {
    (void)json;
}

/* `LDClientCPP` is the v2 C++ binding's RAII client class. Member
 * stubs covered by variadic templates so any arg list (string flag
 * key + fallback, etc.) resolves at parse time. The class is never
 * instantiated; pointers / references are enough. */
class LDClientCPP {
public:
    template <typename... Args> static LDClientCPP *Init(Args&&...) { return nullptr; }
    template <typename... Args> bool boolVariation(Args&&...) { return false; }
    template <typename... Args> bool boolVariationDetail(Args&&...) { return false; }
    template <typename... Args> int intVariationDetail(Args&&...) { return 0; }
    template <typename... Args> double doubleVariationDetail(Args&&...) { return 0; }
    template <typename... Args> int intVariation(Args&&...) { return 0; }
    template <typename... Args> double doubleVariation(Args&&...) { return 0; }
    template <typename... Args> const char *stringVariation(Args&&...) { return ""; }
    LDJSON *getAllFlags() { return nullptr; }
    template <typename... Args> void identify(Args&&...) {}
    template <typename... Args> void track(Args&&...) {}
    template <typename... Args> void close(Args&&...) {}
};

#endif /* LAUNCHDARKLY_API_HPP */
