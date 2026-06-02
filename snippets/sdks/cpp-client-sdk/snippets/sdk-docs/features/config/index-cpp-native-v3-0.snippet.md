---
id: cpp-client-sdk/sdk-docs/features/config/index-cpp-native-v3-0
sdk: cpp-client-sdk
kind: reference
lang: cpp
description: SDK configuration example for C++ (client-side).
---

```cpp
auto config_builder = client_side::ConfigBuilder("example-mobile-key");
config_builder.Events()
    .Capacity(1000)
    .FlushInterval(std::chrono::seconds(30));
auto config = config_builder.Build();
if (!config) {
    /* an error occurred, config is not valid */
}
```
