---
id: cpp-client-sdk/scaffolds/cpp-client-syntax-only-toplevel
sdk: cpp-client-sdk
kind: scaffold
lang: cpp
file: main.cpp
description: |
  File-scope sibling of `cpp-client-syntax-only`. That scaffold
  splices the body inside a nested block of a never-instantiated
  function template, which breaks for fragments that are themselves
  top-level declarations — C++ has no local free functions, so a
  C-binding callback definition like
  `void OnFlagChange(char const* flag_key, ...) { ... }` cannot live
  there. This variant splices the body at file scope instead.

  Same `cpp-client` validator, so the body compiles against the real
  SDK headers from the pre-cloned cpp-sdks checkout.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, spliced at file scope.
validation:
  runtime: cpp-client
  entrypoint: main.cpp
---

```cpp
#include <cstdio>
#include <iostream>
// C-binding headers — top-level callback-definition fragments
// reference LDValue and friends.
#include <launchdarkly/client_side/bindings/c/sdk.h>
#include <launchdarkly/bindings/c/value.h>

{{ body }}

int main() {
    std::cout << "feature flag evaluates to true" << std::endl;
    return 0;
}
```
