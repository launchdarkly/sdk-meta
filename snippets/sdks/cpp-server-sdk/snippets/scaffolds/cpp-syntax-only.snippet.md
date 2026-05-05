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
#include <iostream>
#include <launchdarkly/server_side/client.hpp>
#include <launchdarkly/server_side/bindings/c/sdk.h>

// Wrappee is a never-instantiated template — body is parsed but
// most type-checks are deferred to instantiation (which never
// happens). Keeps doc fragments that reference an undeclared
// `client` (e.g. evaluate-a-context fragments) from failing the
// build.
template <int = 0>
void _wrappee() {
{{ body }}
}

int main() {
    std::cout << "feature flag evaluates to true" << std::endl;
    return 0;
}
```
