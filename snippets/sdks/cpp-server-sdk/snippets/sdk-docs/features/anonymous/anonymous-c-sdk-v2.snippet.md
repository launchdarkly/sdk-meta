---
id: cpp-server-sdk/sdk-docs/features/anonymous/anonymous-c-sdk-v2
sdk: cpp-server-sdk
kind: reference
lang: c
description: Anonymous user example for the C server SDK v2.x.
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only-v2-c
---

```c
struct LDUser *user = LDUserNew("placeholder-key");
LDUserSetAnonymous(user, LDBooleanTrue);
```
