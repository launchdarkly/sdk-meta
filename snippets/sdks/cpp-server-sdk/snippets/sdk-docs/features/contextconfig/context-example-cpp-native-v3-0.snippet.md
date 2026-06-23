---
id: cpp-server-sdk/sdk-docs/features/contextconfig/context-example-cpp-native-v3-0
sdk: cpp-server-sdk
kind: reference
lang: cpp
description: Context example for C++ (server-side) SDK v3.0 (native).
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only

---

```cpp
auto context = ContextBuilder()
  .Kind("user", "example-user-key")
  .Set("firstName", "Sandy")
  .Set("lastName", "Smith")
  .Set("groups", {"Acme", "Global Health Services"})
  .Build();
```
