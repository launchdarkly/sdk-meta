---
id: cpp-server-sdk/sdk-docs/migration-2-to-3-understanding-differences-between-users-and-contexts-3-0-syntax-single-context-with-key
sdk: cpp-server-sdk
kind: reference
lang: cpp
description: "3.0 syntax, single context with key in section \"Understanding differences between users and contexts\""
---

```cpp
auto context = ContextBuilder()
  .Kind("organization", "example-organization-key")
  .Build();
```
