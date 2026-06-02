---
id: cpp-client-sdk/sdk-docs/features/config/app-config-cpp-native-v3-0
sdk: cpp-client-sdk
kind: reference
lang: cpp
description: Application metadata configuration example for C++ (client-side).
---

```cpp
auto config_builder = client_side::ConfigBuilder("example-mobile-key");
config_builder.AppInfo().Identifier("authentication-service").Version("1.0.0")
auto config = config_builder.Build();
```
