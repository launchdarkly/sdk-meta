---
id: cpp-server-sdk/sdk-docs/migration-2-to-3-understanding-changes-to-private-attributes-3-0-syntax-all-attributes-marked-private
sdk: cpp-server-sdk
kind: reference
lang: cpp
description: "3.0 syntax, all attributes marked private in section \"Understanding changes to private attributes\""
---

```cpp
auto config_builder = server_side::ConfigBuilder("YOUR_SDK_KEY");
config_builder.Events().AllAttributesPrivate(true);
auto config = config_builder.Build();
```
