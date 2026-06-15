---
id: cpp-server-sdk/scaffolds/cpp-syntax-only
sdk: cpp-server-sdk
kind: scaffold
lang: cpp
file: main.cpp
description: |
  Parse-only validator for C++ server SDK doc fragments.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: cpp-server
  entrypoint: main.cpp
---

```cpp
#include <chrono>
#include <cstdio>
#include <future>
#include <iostream>
#include <optional>
#include <string>
// Native C++ headers.
#include <launchdarkly/server_side/client.hpp>
#include <launchdarkly/server_side/config/config_builder.hpp>
#include <launchdarkly/context_builder.hpp>
#include <launchdarkly/value.hpp>
// C-binding headers — doc fragments mix C-binding and native styles.
#include <launchdarkly/server_side/bindings/c/sdk.h>
#include <launchdarkly/server_side/bindings/c/config/builder.h>
#include <launchdarkly/bindings/c/context.h>
#include <launchdarkly/bindings/c/context_builder.h>

// Polymorphic stub so a body can use `client.BoolVariation(...)`
// (native-style) AND `LDServerSDK_BoolVariation(client, ...)`
// (C-binding-style) without needing the scaffold to know which
// shape the body uses. The wrappee is a never-instantiated template,
// so the conversion operator and member functions exist at the type
// system level but are never invoked.
// Stub of the data-source-status provider returned by the native
// API's `client.DataSourceStatus()` — variadic member so the body's
// `.OnDataSourceStatusChange(lambda)` chain type-checks (the lambda
// itself is checked against the real DataSourceStatus type).
struct _AnyStatusProvider {
    template <typename... Args> int OnDataSourceStatusChange(Args&&...) const { return 0; }
};

// Stub matching the C-binding listener callback the docs define in a
// separate code block; bodies that only show the assignment
// (`listener.StatusChanged = OnDataSourceStatusChanged;`) resolve
// against this. Signature mirrors ServerDataSourceStatusCallbackFn.
inline void OnDataSourceStatusChanged(LDServerDataSourceStatus status, void* user_data) {
    (void)status;
    (void)user_data;
}

struct _AnyClient {
    operator LDServerSDK() const { return nullptr; }
    // operator-> makes `client->Method(...)` resolve when the body
    // uses pointer syntax (typical of the C-binding API where
    // `LDServerSDK` is an opaque pointer alias). Self-pointer so
    // the chained call dispatches to the same variadic stubs below.
    const _AnyClient* operator->() const { return this; }
    template <typename... Args> bool BoolVariation(Args&&...) const { return false; }
    // Detail-variation stub returning a real EvaluationDetail so the
    // body's `detail.Value()` / `detail.Reason()` chains type-check.
    template <typename... Args> auto BoolVariationDetail(Args&&...) const {
        return launchdarkly::EvaluationDetail<bool>(false, std::nullopt, std::nullopt);
    }
    template <typename... Args> int IntVariation(Args&&...) const { return 0; }
    template <typename... Args> double DoubleVariation(Args&&...) const { return 0; }
    template <typename... Args> std::string StringVariation(Args&&...) const { return {}; }
    template <typename... Args> auto JsonVariation(Args&&...) const { return launchdarkly::Value{}; }
    // The real client's AllFlagsState returns the AllFlagsState class;
    // the stub returns a default-constructed (invalid) instance so the
    // body's `.Valid()` / `.Values()` chains compile.
    template <typename... Args> auto AllFlagsState(Args&&...) const {
        return launchdarkly::server_side::AllFlagsState{};
    }
    template <typename... Args> void TrackEvent(Args&&...) const {}
    template <typename... Args> void Identify(Args&&...) const {}
    template <typename... Args> auto StartAsync(Args&&...) const { return std::async(std::launch::deferred, []{ return false; }); }
    template <typename... Args> _AnyStatusProvider DataSourceStatus(Args&&...) const { return {}; }
};

// Wrappee is a never-instantiated template — body is parsed but
// most type-checks are deferred to instantiation (which never
// happens). The body lives in a nested block so it can re-declare
// `client` / `context` when the fragment shows the native or
// C-binding init form (e.g. `server_side::Client client(*config);`,
// `LDServerSDK client = LDServerSDK_New(...)`).
template <int = 0>
void _wrappee() {
    // Doc fragments mix unqualified names from `launchdarkly` and
    // `launchdarkly::server_side` (e.g. `ContextBuilder`,
    // `ConfigBuilder`, `Client`, `server_side::ConfigBuilder`).
    // Lifting both namespaces inside the template body lets bodies
    // that omit the namespace prefix (the "no-namespace" /
    // "using-launchdarkly-server-side-namespace" doc style) compile
    // without losing the namespace-qualified bodies — those still
    // resolve through the qualified name first.
    using namespace launchdarkly;
    using namespace launchdarkly::server_side;
    _AnyClient client;
    LDContext context = nullptr;
    LDServerConfig config = nullptr;
    // The listener fragments split "create the connection" and "free
    // the connection" across separate code blocks; the free-side body
    // references `connection` as if pre-existing.
    LDListenerConnection connection = nullptr;
    // `maxwait` is referenced by both native-style fragments
    // (`wait_for(maxwait)` — needs a chrono duration) and C-binding
    // fragments (`LDServerSDK_Start(client, maxwait, ...)` — needs an
    // unsigned int milliseconds count). The polymorphic stub provides
    // implicit conversions to both so neither shape needs to declare
    // it locally.
    struct _Maxwait {
        operator unsigned int() const { return 10000; }
        operator std::chrono::milliseconds() const { return std::chrono::milliseconds{10000}; }
        operator std::chrono::seconds() const { return std::chrono::seconds{10}; }
    } maxwait;
    {
{{ body }}
    }
}

int main() {
    std::cout << "feature flag evaluates to true" << std::endl;
    return 0;
}
```
