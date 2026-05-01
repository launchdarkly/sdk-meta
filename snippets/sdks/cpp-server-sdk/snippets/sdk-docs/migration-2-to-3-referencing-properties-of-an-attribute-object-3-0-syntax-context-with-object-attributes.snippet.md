---
id: cpp-server-sdk/sdk-docs/migration-2-to-3-referencing-properties-of-an-attribute-object-3-0-syntax-context-with-object-attributes
sdk: cpp-server-sdk
kind: reference
lang: cpp
description: "3.0 syntax, context with object attributes in section \"Referencing properties of an attribute object\""
---

```cpp
auto context = ContextBuilder()
  .Kind("user", "example-user-key")
  .Set("address", Value::Object({{"street", "123 Main St"}, {"city", "Springfield"}}))
  .Build();
```
