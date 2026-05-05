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
  themselves or the build catches the issue. v2.x fragments are
  Bucket C — see `_sdk-docs-port-notes.md`.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: cpp-client
  entrypoint: main.cpp
---

```cpp
#include <iostream>
#include <launchdarkly/client_side/client.hpp>
#include <launchdarkly/client_side/bindings/c/sdk.h>

template <int = 0>
void _wrappee() {
{{ body }}
}

int main() {
    std::cout << "feature flag evaluates to true" << std::endl;
    return 0;
}
```
