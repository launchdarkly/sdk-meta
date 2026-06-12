---
id: cpp-client-sdk/scaffolds/cpp-client-syntax-only-toplevel
sdk: cpp-client-sdk
kind: scaffold
lang: cpp
file: main.cpp
description: |
  Parse-only validator for C++ client SDK doc fragments whose shape
  is a top-level declaration (e.g. a callback function definition for
  the C-binding listener API). C++ forbids function definitions
  inside another function, so these fragments cannot ride the
  default cpp-client-syntax-only scaffold's `_wrappee()` splice; the
  body is spliced at file scope instead. Same include set as the
  default scaffold so both native and C-binding names resolve.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, spliced at file scope.
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
// Native C++ headers.
#include <launchdarkly/client_side/client.hpp>
#include <launchdarkly/context_builder.hpp>
#include <launchdarkly/value.hpp>
// C-binding headers — doc fragments mix C-binding and native styles.
#include <launchdarkly/client_side/bindings/c/sdk.h>
#include <launchdarkly/client_side/bindings/c/config/builder.h>
#include <launchdarkly/bindings/c/context.h>
#include <launchdarkly/bindings/c/context_builder.h>

// The fragment shows `printf` without its own include; <cstdio>
// above provides it in the global namespace on the toolchains the
// validator uses.
{{ body }}

int main() {
    std::cout << "feature flag evaluates to true" << std::endl;
    return 0;
}
```
