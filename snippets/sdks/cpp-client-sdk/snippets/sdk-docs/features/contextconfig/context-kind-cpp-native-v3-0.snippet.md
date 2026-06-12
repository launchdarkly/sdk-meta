---
id: cpp-client-sdk/sdk-docs/features/contextconfig/context-kind-cpp-native-v3-0
sdk: cpp-client-sdk
kind: reference
lang: cpp
description: Context with a non-user kind for C++ (client-side) SDK v3.0 (native).
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only

---

```cpp
auto context = ContextBuilder()
  .Kind("organization", "example-organization-key")
  .Build();
```
