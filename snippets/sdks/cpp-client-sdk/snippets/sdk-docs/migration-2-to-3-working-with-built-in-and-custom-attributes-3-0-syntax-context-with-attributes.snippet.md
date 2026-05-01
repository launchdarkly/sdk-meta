---
id: cpp-client-sdk/sdk-docs/migration-2-to-3-working-with-built-in-and-custom-attributes-3-0-syntax-context-with-attributes
sdk: cpp-client-sdk
kind: reference
lang: cpp
description: "3.0 syntax, context with attributes in section \"Working with built-in and custom attributes\""
---

```cpp
auto context = ContextBuilder()
  .Kind("user", "example-user-key")
  .Name("Sandy")
  .Set("email", "sandy@example.com")
  .Build();
```
