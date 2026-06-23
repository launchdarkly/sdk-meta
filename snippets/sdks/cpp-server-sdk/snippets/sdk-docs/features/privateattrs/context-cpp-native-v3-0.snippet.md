---
id: cpp-server-sdk/sdk-docs/features/privateattrs/context-cpp-native-v3-0
sdk: cpp-server-sdk
kind: reference
lang: cpp
description: Marking context attributes private with the context builder for C++ server SDK v3.0.
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only

---

```cpp
auto context = ContextBuilder()
  .Kind("user", "example-user-key")
  .SetPrivate("email", "sandy@example.com")
  .Build();
```
