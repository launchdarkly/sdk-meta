---
id: cpp-server-sdk/sdk-docs/features/contextconfig/multi-context-cpp-native-v3-0
sdk: cpp-server-sdk
kind: reference
lang: cpp
description: Multi-context example for C++ (server-side) SDK v3.0 (native).
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only

---

```cpp
auto context = ContextBuilder()
  .Kind("user", "example-user-key")
  .Name("Sandy")
  .Kind("organization", "example-organization-key")
  .Name("Global Health Services")
  .Build();
```
