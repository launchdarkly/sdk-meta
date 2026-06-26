---
id: cpp-server-sdk/sdk-docs/features/anonymous/anonymous-cpp-native-v3-0
sdk: cpp-server-sdk
kind: reference
lang: cpp
description: Anonymous context example for C++ (server-side), native API.
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only
---

```cpp
auto context = ContextBuilder()
  .Kind("user", "example-user-key")
  .Anonymous(true)
  .Build();
```
