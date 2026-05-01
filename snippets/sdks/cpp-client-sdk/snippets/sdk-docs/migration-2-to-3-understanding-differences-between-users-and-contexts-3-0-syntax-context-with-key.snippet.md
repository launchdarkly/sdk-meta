---
id: cpp-client-sdk/sdk-docs/migration-2-to-3-understanding-differences-between-users-and-contexts-3-0-syntax-context-with-key
sdk: cpp-client-sdk
kind: reference
lang: cpp
description: "3.0 syntax, context with key in section \"Understanding differences between users and contexts\""
---

```cpp
auto context = ContextBuilder()
  .Kind("user", "example-user-key")
  .Build();
```
