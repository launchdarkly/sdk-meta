---
id: cpp-client-sdk/sdk-docs/features/contextconfig/context-example-cpp-native-v3-0
sdk: cpp-client-sdk
kind: reference
lang: cpp
description: Context example for C++ (client-side) SDK v3.0 (native).
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only

---

```cpp
auto context = ContextBuilder()
  .Kind("user", "example-user-key")
  .Set("firstName", "Sandy")
  .Set("lastName", "Smith")
  .Set("groups", {"Acme", "Global Health Services"})
  .Build();
```
