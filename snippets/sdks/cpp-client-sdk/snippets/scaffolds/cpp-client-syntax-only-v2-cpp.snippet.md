---
id: cpp-client-sdk/scaffolds/cpp-client-syntax-only-v2-cpp
sdk: cpp-client-sdk
kind: scaffold
lang: cpp
file: snippet.cpp
description: |
  Parse-and-type-check validator for legacy v2.x cpp-client SDK doc
  fragments that target the C++ binding API surface (LDClientCPP,
  `client->boolVariation(...)`, `<launchdarkly/api.hpp>`).

  Routes through the `cpp-client-v2-cpp` validator container, which
  ships a stub `<launchdarkly/api.hpp>` declaring the v2 C++
  binding's `LDClientCPP` class and the C-flavored helpers
  (`LDConfigNew`, `LDUserNew`) that v2 doc fragments use. The
  current cpp-client validator builds against cpp-sdks v3+, where
  `api.hpp` and `LDClientCPP` are gone.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: cpp-client-v2-cpp
  entrypoint: snippet.cpp
---

```cpp
#include <launchdarkly/api.hpp>

// File-scope stubs for fragments that read like statement bodies
// (`bool show_feature = client->boolVariation(…)` — referencing a
// pre-declared `client`). The doc assumes the reader's earlier init
// snippets already declared these.
static LDClientCPP *client;
static LDUser *user;
static LDConfig *config;
static unsigned int maxwait;

static void _wrappee() {
{{ body }}
}

int main() {
    (void)_wrappee;
    return 0;
}
```
