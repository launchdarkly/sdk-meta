---
id: cpp-client-sdk/scaffolds/cpp-client-syntax-only
sdk: cpp-client-sdk
kind: scaffold
lang: cpp
file: main.cpp
description: |
  Parse-only validator for C++ client SDK doc fragments.

  Wrappee is a never-instantiated template so the body is
  syntactically parsed but type-checks against unbound names are
  deferred until instantiation (which never happens). Include the
  SDK headers at file scope so doc fragments that show
  `#include <launchdarkly/...>` directives at the top of the body
  re-include cheaply (header guards) and without conflict.

  Doc fragments that don't declare a `client` or `context` and
  reference them as if pre-existing won't compile through this
  scaffold; in practice the v3 fragments either declare those
  themselves or the build catches the issue. v2.x fragments may
  need additional stubs on `_AnyClient` — see
  `_sdk-docs-port-notes.md`.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: cpp-client
  entrypoint: main.cpp
---

```cpp
#include <chrono>
#include <cstdio>
#include <future>
#include <iostream>
#include <optional>
#include <string>
#include <unordered_map>
// Native C++ headers.
#include <launchdarkly/client_side/client.hpp>
#include <launchdarkly/context_builder.hpp>
#include <launchdarkly/value.hpp>
// C-binding headers — doc fragments mix C-binding and native styles.
#include <launchdarkly/client_side/bindings/c/sdk.h>
#include <launchdarkly/client_side/bindings/c/config/builder.h>
#include <launchdarkly/bindings/c/context.h>
#include <launchdarkly/bindings/c/context_builder.h>
// array_builder.h is included at file scope so doc fragments that
// carry their own in-body `#include <.../array_builder.h>` line
// hit the header's include guard there (a first include inside the
// wrappee's function body would be invalid C++ — the header opens
// an extern "C" block).
#include <launchdarkly/bindings/c/array_builder.h>

// Polymorphic stub so a body can use `client.BoolVariation(...)`
// (native-style) AND `LDClientSDK_BoolVariation(client, ...)`
// (C-binding-style) without needing the scaffold to know which
// shape the body uses. The wrappee is a never-instantiated template,
// so the conversion operator and member functions exist at the type
// system level but are never invoked.
struct _AnyClient {
    operator LDClientSDK() const { return nullptr; }
    // operator-> makes `client->Method(...)` resolve when the body
    // uses pointer syntax (typical of the C-binding API where
    // `LDClientSDK` is an opaque pointer alias). The arrow returns
    // a self-pointer so the chained call dispatches to the same
    // variadic stubs below.
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
    // The real client's AllFlags returns a flag-key-to-Value map; the
    // stub matches so range-for fragments with structured bindings
    // compile.
    template <typename... Args> auto AllFlags(Args&&...) const {
        return std::unordered_map<std::string, launchdarkly::Value>{};
    }
    template <typename... Args> void Track(Args&&...) const {}
    template <typename... Args> void TrackEvent(Args&&...) const {}
    template <typename... Args> void Identify(Args&&...) const {}
    // Identify-and-examine-the-result fragments treat the return as a
    // future<bool>, mirroring the real v3 client's IdentifyAsync.
    template <typename... Args> auto IdentifyAsync(Args&&...) const { return std::async(std::launch::deferred, []{ return false; }); }
    // Matches the real Client::FlushAsync surface: fire-and-forget,
    // returns void.
    template <typename... Args> void FlushAsync(Args&&...) const {}
    template <typename... Args> auto StartAsync(Args&&...) const { return std::async(std::launch::deferred, []{ return false; }); }
    // Lowercase-first aliases — the v2.x C++ client SDK exposed
    // camelCased methods (e.g. `client->boolVariation(...)`); v3.x
    // renamed to PascalCase. Doc fragments still cover both eras, so
    // expose both surfaces.
    template <typename... Args> bool boolVariation(Args&&...) const { return false; }
    template <typename... Args> int intVariation(Args&&...) const { return 0; }
    template <typename... Args> double doubleVariation(Args&&...) const { return 0; }
    template <typename... Args> std::string stringVariation(Args&&...) const { return {}; }
    template <typename... Args> auto jsonVariation(Args&&...) const { return launchdarkly::Value{}; }
    template <typename... Args> auto allFlags(Args&&...) const { return launchdarkly::Value{}; }
    template <typename... Args> void trackEvent(Args&&...) const {}
    template <typename... Args> void identify(Args&&...) const {}
    template <typename... Args> auto startAsync(Args&&...) const { return std::async(std::launch::deferred, []{ return false; }); }
};

// Polymorphic stub for the ambient `config_builder` some doc
// fragments reference (the docs assume an earlier init fragment
// declared it). Satisfies both the native member-call shape
// (`config_builder.Offline(true)`) and the C-binding shape
// (`LDClientConfigBuilder_Offline(config_builder, true)`) via an
// implicit conversion to the opaque builder handle. File-scope
// because local classes cannot declare member templates.
struct _AnyConfigBuilder {
    operator LDClientConfigBuilder() const { return nullptr; }
    const _AnyConfigBuilder* operator->() const { return this; }
    template <typename... Args> void Offline(Args&&...) const {}
};

template <int = 0>
void _wrappee() {
    // Body lives in a nested block so it can re-declare `client` /
    // `context` (e.g. native-API init fragments that say
    // `client_side::Client client(...)`, or C-binding init fragments
    // that say `LDClientSDK client = LDClientSDK_New(...)`) without
    // colliding with the stubs above. Bodies that use `client`
    // without declaring it (e.g. evaluate-a-flag fragments) pick up
    // the polymorphic stub from the outer scope.
    //
    // Doc fragments mix unqualified `ContextBuilder`,
    // `client_side::Client`, etc. with the
    // `launchdarkly::` prefix. Lifting both namespaces lets bodies
    // that omit the prefix compile without breaking
    // namespace-qualified bodies.
    using namespace launchdarkly;
    using namespace launchdarkly::client_side;
    _AnyClient client;
    _AnyConfigBuilder config_builder;
    LDContext context = nullptr;
    // Identify fragments pass an `updated_context` built by an earlier
    // fragment; the docs assume it already exists.
    LDContext updated_context = nullptr;
    LDClientConfig config = nullptr;
    // `maxwait` is referenced by both native-style fragments
    // (`wait_for(maxwait)` — needs a chrono duration) and C-binding
    // fragments (`LDClientSDK_Start(client, maxwait, ...)` — needs an
    // unsigned int milliseconds count). The polymorphic stub provides
    // implicit conversions to both.
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
