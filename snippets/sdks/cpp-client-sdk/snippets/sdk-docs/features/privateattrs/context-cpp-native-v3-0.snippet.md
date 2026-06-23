---
id: cpp-client-sdk/sdk-docs/features/privateattrs/context-cpp-native-v3-0
sdk: cpp-client-sdk
kind: reference
lang: cpp
description: Marking context attributes private with the context builder for C++ client SDK v3.0.
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only

---

```cpp
auto context = ContextBuilder()
  .Kind("user", "example-user-key")
  .Name("Sandy Smith")
  .SetPrivate("email", "sandy@example.com")
  .Build();
```
