---
id: cpp-server-sdk/scaffolds/cpp-syntax-only-redis
sdk: cpp-server-sdk
kind: scaffold
lang: cpp
file: main.cpp
description: |
  Parse-only validator for C++ server SDK doc fragments that reference
  the Redis integration (`launchdarkly::server_redis_source`).

  Same shape as `cpp-syntax-only`, with two additions:

  - The Redis source headers (native + C binding) are pre-included at
    file scope. Doc fragments carry their own `#include` of
    `redis_source.hpp` inside the wrappee body; the header's include
    guard makes that inner include a no-op, so the body stays verbatim
    while the declarations land at file scope where they belong.
  - `validation.env` sets `CPP_REDIS=1`, which tells the cpp-server
    harness to configure cpp-sdks with `-DLD_BUILD_REDIS_SUPPORT=ON`
    and link `launchdarkly::server_redis_source`.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: cpp-server
  entrypoint: main.cpp
  env:
    CPP_REDIS: "1"
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
// Redis source headers (native + C binding). Bodies that #include the
// native header themselves hit the include guard and stay verbatim.
#include <launchdarkly/server_side/integrations/redis/redis_source.hpp>
#include <launchdarkly/server_side/bindings/c/integrations/redis/redis_source.h>

// Wrappee is a never-instantiated template — body is parsed but
// most type-checks are deferred to instantiation (which never
// happens). The body lives in a nested block so it can re-declare
// names the scaffold stubs.
template <int = 0>
void _wrappee() {
    // Unlike cpp-syntax-only, no namespaces are lifted here: the
    // native Redis fragment carries its own
    // `using namespace launchdarkly::server_side;` and refers to
    // `config::builders::...` unqualified — lifting `launchdarkly`
    // as well would make `config` ambiguous
    // (launchdarkly::config vs launchdarkly::server_side::config).
    // Init-shaped fragments pass an `sdk_key` the docs assume already
    // exists.
    char const* sdk_key = "";
    (void)sdk_key;
    {
{{ body }}
    }
}

int main() {
    std::cout << "feature flag evaluates to true" << std::endl;
    return 0;
}
```
