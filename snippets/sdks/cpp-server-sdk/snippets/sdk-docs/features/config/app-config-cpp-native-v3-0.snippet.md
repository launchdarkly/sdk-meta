---
id: cpp-server-sdk/sdk-docs/features/config/app-config-cpp-native-v3-0
sdk: cpp-server-sdk
kind: reference
lang: cpp
description: Application metadata configuration example for C++ (server-side).
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only

---

```cpp
auto config_builder = server_side::ConfigBuilder("YOUR_SDK_KEY");
config_builder.AppInfo().Identifier("authentication-service").Version("1.0.0");
auto config = config_builder.Build();
if (!config) {
    /* an error occurred, config is not valid */
}
```
