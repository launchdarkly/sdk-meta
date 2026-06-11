---
id: cpp-server-sdk/scaffolds/cpp-syntax-only-v2-cpp
sdk: cpp-server-sdk
kind: scaffold
lang: cpp
file: snippet.cpp
description: |
  Parse-and-type-check validator for C++-flavored doc fragments of
  the legacy v2.x C server SDK. The v2 server SDK had no separate C++
  binding class — its "C++ binding" doc fragments call the C API from
  C++ — so this scaffold mirrors `cpp-syntax-only-v2-c` but routes
  through the `cpp-server-v2-cpp` validator, whose g++ pass accepts
  the fragments' C++-isms (`bool`, unqualified struct names).
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: cpp-server-v2-cpp
  entrypoint: snippet.cpp
---

```cpp
#include <launchdarkly/api.h>

/* File-scope stubs so fragments that read like statement bodies
 * resolve at compile time. The doc fragments assume the reader's
 * prior init snippets already declared these. */
static struct LDClient *client;
static struct LDUser *user;
static struct LDConfig *config;
static unsigned int maxwaitmilliseconds;

static void _wrappee(void) {
{{ body }}
}

int main(void) {
    (void)_wrappee;
    return 0;
}
```
