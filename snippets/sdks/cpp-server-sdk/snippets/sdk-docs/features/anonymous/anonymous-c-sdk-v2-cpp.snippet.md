---
id: cpp-server-sdk/sdk-docs/features/anonymous/anonymous-c-sdk-v2-cpp
sdk: cpp-server-sdk
kind: reference
lang: cpp
description: Anonymous user example for the C server SDK v2.x, C++ binding.
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only-v2-cpp
---

```cpp
struct LDUser *user = LDUserNew("placeholder-key");
LDUserSetAnonymous(user, true);
```
