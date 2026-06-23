---
id: cpp-server-sdk/scaffolds/cpp-syntax-only-v2-c-toplevel
sdk: cpp-server-sdk
kind: scaffold
lang: c
file: snippet.c
description: |
  File-scope variant of `cpp-syntax-only-v2-c` for legacy v2.x C
  server SDK doc fragments that are themselves top-level declarations
  (e.g. a `static` custom logger function definition). C forbids
  `static` function definitions inside another function, so these
  fragments cannot route through the nested `_wrappee()` scaffold.

  The body is spliced at file scope after the stub
  `<launchdarkly/api.h>` header; `<stdio.h>` is included for
  fragments that call `printf`. Nothing in the body is ever invoked.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, spliced at file scope.
validation:
  runtime: cpp-server-v2-c
  entrypoint: snippet.c
---

```c
#include <stdio.h>

#include <launchdarkly/api.h>

{{ body }}

int main(void) {
    return 0;
}
```
