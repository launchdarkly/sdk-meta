---
id: cpp-client-sdk/sdk-docs/migration-2-to-3-understanding-changes-to-private-attributes-3-0-syntax-attribute-marked-private-for-one-context
sdk: cpp-client-sdk
kind: reference
lang: cpp
description: "3.0 syntax, attribute marked private for one context in section \"Understanding changes to private attributes\""
---

```cpp
auto context = ContextBuilder()
  .Kind("user", "example-user-key")
  .Name("Sandy Smith")
  .SetPrivate("email", "sandy@example.com")
  .Build();
```
