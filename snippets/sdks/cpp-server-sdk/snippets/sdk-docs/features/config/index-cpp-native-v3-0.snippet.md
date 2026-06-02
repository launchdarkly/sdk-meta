---
id: cpp-server-sdk/sdk-docs/features/config/index-cpp-native-v3-0
sdk: cpp-server-sdk
kind: reference
lang: cpp
description: SDK configuration example for C++ (server-side).
---

```cpp
auto config_builder = server_side::ConfigBuilder("YOUR_SDK_KEY");
config_builder.Events()
    .Capacity(1000)
    .FlushInterval(std::chrono::seconds(30));
auto config = config_builder.Build();
if (!config) {
    /* an error occurred, config is not valid */
}
```
