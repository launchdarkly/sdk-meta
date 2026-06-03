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

/* C-flavored types live in the global namespace, matching the
 * v2 SDK's actual layout (the C and C++ headers shared a translation
 * unit). */
typedef bool LDBoolean;

struct LDConfig;
struct LDUser;

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

/* `LDClientCPP` is the v2 C++ binding's RAII client class. Member
 * stubs covered by variadic templates so any arg list (string flag
 * key + fallback, etc.) resolves at parse time. The class is never
 * instantiated; pointers / references are enough. */
class LDClientCPP {
public:
    template <typename... Args> static LDClientCPP *Init(Args&&...) { return nullptr; }
    template <typename... Args> bool boolVariation(Args&&...) { return false; }
    template <typename... Args> int intVariation(Args&&...) { return 0; }
    template <typename... Args> double doubleVariation(Args&&...) { return 0; }
    template <typename... Args> const char *stringVariation(Args&&...) { return ""; }
    template <typename... Args> void identify(Args&&...) {}
    template <typename... Args> void track(Args&&...) {}
    template <typename... Args> void close(Args&&...) {}
};

#endif /* LAUNCHDARKLY_API_HPP */
