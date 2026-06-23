---
id: cpp-client-sdk/sdk-docs/features/privateattrs/config-cpp-native-v3-0
sdk: cpp-client-sdk
kind: reference
lang: cpp
description: Private attribute configuration for C++ client SDK v3.0 (native).
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only

---

```cpp
/* sets all attributes private */
auto config_builder = client_side::ConfigBuilder("example-mobile-key");
config_builder.Events().AllAttributesPrivate(true);
auto config_all_private = config_builder.Build();

/* sets "email" and "address" private */
config_builder = client_side::ConfigBuilder("example-mobile-key");
config_builder.Events().PrivateAttributes({"email", "address"});
auto configSomePrivate = config_builder.Build();
```
