---
id: cpp-client-sdk/sdk-docs/features/anonymous/anonymous-cpp-native-v3-0
sdk: cpp-client-sdk
kind: reference
lang: cpp
description: Anonymous context example for C++ (client-side), native API.
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only
---

```cpp
auto context = ContextBuilder()
  .Kind("user", "example-user-key")
  .Anonymous(true)
  .Build();
```
