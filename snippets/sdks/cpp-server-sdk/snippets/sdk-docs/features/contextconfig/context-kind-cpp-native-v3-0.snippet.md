---
id: cpp-server-sdk/sdk-docs/features/contextconfig/context-kind-cpp-native-v3-0
sdk: cpp-server-sdk
kind: reference
lang: cpp
description: Context with a non-user kind for C++ (server-side) SDK v3.0 (native).
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only

---

```cpp
auto context = ContextBuilder()
  .Kind("organization", "example-organization-key")
  .Build();
```
