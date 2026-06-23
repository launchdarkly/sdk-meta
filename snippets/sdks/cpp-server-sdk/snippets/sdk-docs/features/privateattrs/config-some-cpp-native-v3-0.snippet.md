---
id: cpp-server-sdk/sdk-docs/features/privateattrs/config-some-cpp-native-v3-0
sdk: cpp-server-sdk
kind: reference
lang: cpp
description: Marking specific attributes private for C++ server SDK v3.0 (native).
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only

---

```cpp
auto config_builder = server_side::ConfigBuilder("YOUR_SDK_KEY");
config_builder.Events().PrivateAttributes({"email"});
auto config = config_builder.Build();
if (!config) {
    /* an error occurred, config is not valid */
}
```
