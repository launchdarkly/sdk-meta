---
id: cpp-client-sdk/sdk-docs/migration-2-to-3-understanding-changes-to-private-attributes-3-0-syntax-all-attributes-marked-private
sdk: cpp-client-sdk
kind: reference
lang: cpp
description: "3.0 syntax, all attributes marked private in section \"Understanding changes to private attributes\""
---

```cpp
auto config_builder = client_side::ConfigBuilder("example-mobile-key");
config_builder.Events().AllAttributesPrivate(true);
auto config = config_builder.Build();
```
