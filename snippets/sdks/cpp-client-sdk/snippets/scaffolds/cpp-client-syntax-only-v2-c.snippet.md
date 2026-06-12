---
id: cpp-client-sdk/scaffolds/cpp-client-syntax-only-v2-c
sdk: cpp-client-sdk
kind: scaffold
lang: c
file: snippet.c
description: |
  Parse-and-type-check validator for legacy v2.x C client SDK doc
  fragments (the `launchdarkly/c-client-sdk` package, distinct from
  `c-server-sdk`).

  Routes through the `cpp-client-v2-c` validator container, which
  ships a minimal stub of `<launchdarkly/api.h>` declaring the v2
  client-flavored API (LDClientInit takes config + user + maxwait;
  LDBoolVariation takes client + flagKey + fallback — both shapes
  differ from the v2 server SDK's).
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: cpp-client-v2-c
  entrypoint: snippet.c
---

```c
#include <launchdarkly/api.h>

/* File-scope stubs for fragments that read like statement bodies
 * (`bool show_feature = LDBoolVariation(client, …)` — referencing
 * a pre-declared `client`). The doc assumes the reader's earlier
 * init snippets already declared these. */
static struct LDClient *client;
static struct LDUser *user;
static struct LDConfig *config;
static unsigned int maxwait;
static unsigned int maxwaitmilliseconds;
static struct LDUser *newUser;
static struct LDUser *previousUser;

static void _wrappee(void) {
{{ body }}
}

int main(void) {
    (void)_wrappee;
    return 0;
}
```
