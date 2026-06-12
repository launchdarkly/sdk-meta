---
id: cpp-server-sdk/scaffolds/cpp-syntax-only-toplevel
sdk: cpp-server-sdk
kind: scaffold
lang: cpp
file: main.cpp
description: |
  File-scope variant of `cpp-syntax-only` for C++ server SDK doc
  fragments that are themselves top-level declarations -- custom log
  backend classes, C-binding callback function definitions -- which
  cannot live inside the nested-block `_wrappee()` body (C++ forbids
  function definitions inside a function, and `static` storage on a
  local declaration of one is a hard error).

  The body is spliced at file scope after the SDK headers. Fragments
  that carry their own `#include` directives re-include cheaply
  (header guards / pragma once) because the same headers are already
  included here. Nothing in the body is ever invoked; `main()` just
  prints the EXAM-HELLO sentinel.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, spliced at file scope.
validation:
  runtime: cpp-server
  entrypoint: main.cpp
---

```cpp
#include <cstdio>
#include <iostream>
#include <memory>
#include <optional>
#include <string>
// Native C++ headers, including the logging interface custom-backend
// fragments implement.
#include <launchdarkly/server_side/client.hpp>
#include <launchdarkly/server_side/config/config_builder.hpp>
#include <launchdarkly/logging/log_backend.hpp>
#include <launchdarkly/logging/log_level.hpp>
// C-binding headers -- doc fragments mix C-binding and native styles.
#include <launchdarkly/server_side/bindings/c/sdk.h>
#include <launchdarkly/server_side/bindings/c/config/builder.h>

{{ body }}

int main() {
    std::cout << "feature flag evaluates to true" << std::endl;
    return 0;
}
```
