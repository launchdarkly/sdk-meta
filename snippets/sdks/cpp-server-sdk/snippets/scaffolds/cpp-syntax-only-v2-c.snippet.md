---
id: cpp-server-sdk/scaffolds/cpp-syntax-only-v2-c
sdk: cpp-server-sdk
kind: scaffold
lang: c
file: snippet.c
description: |
  Parse-and-type-check validator for legacy v2.x C server SDK doc
  fragments (the `launchdarkly/c-server-sdk` package, removed when
  cpp-server was rewritten on top of the modern cpp-sdks C++
  codebase).

  Routes through the `cpp-server-v2-c` validator container, which
  ships a minimal stub of `<launchdarkly/api.h>` and runs the staged
  source through `gcc -c`. The body lives inside a `_wrappee()`
  function so bodies that read like statement fragments (assigning
  to a pre-declared `user`, etc.) have a function scope to land in.
  File-scope stubs declare the v2-era `client`, `user`, `config`,
  `maxwaitmilliseconds` symbols the docs assume already exist.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: cpp-server-v2-c
  entrypoint: snippet.c
---

```c
#include <launchdarkly/api.h>

/* File-scope stubs so fragments that read like statement bodies
 * (`user = LDUserNew(...);` — assignment to a pre-declared `user`)
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
