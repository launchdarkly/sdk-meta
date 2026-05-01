---
id: cpp-server-sdk/sdk-docs/migration-2-to-3-understanding-changes-to-private-attributes-3-0-syntax-specific-fields-of-an-attribute-marked-private
sdk: cpp-server-sdk
kind: reference
lang: cpp
description: "3.0 syntax, specific fields of an attribute marked private in section \"Understanding changes to private attributes\""
---

```cpp
auto context = ContextBuilder()
  .Kind("user", "example-user-key")
  .Name("Sandy Smith")
  .Set("address", Value::Object({{"street", "Main St"}, {"city", "Springfield"}}))
  .AddPrivateAttribute("/address/street")
  .Build()
```
