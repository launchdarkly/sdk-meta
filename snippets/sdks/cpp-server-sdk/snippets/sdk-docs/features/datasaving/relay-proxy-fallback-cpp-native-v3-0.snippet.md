---
id: cpp-server-sdk/sdk-docs/features/datasaving/relay-proxy-fallback-cpp-native-v3-0
sdk: cpp-server-sdk
kind: reference
lang: cpp
description: Data saving mode with Relay Proxy and LaunchDarkly API fallback for C++ (server-side).
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only
---

```cpp
using FDv2 = server_side::config::builders::DataSystemBuilder::FDv2;

std::string relay_url = "http://my-relay-proxy:8030";

auto config_builder = server_side::ConfigBuilder("YOUR_SDK_KEY");
config_builder.DataSystem().Method(
    FDv2::Custom()
        .Initializer(FDv2::Polling().BaseUrl(relay_url))
        .Initializer(FDv2::Polling())
        .Synchronizer(FDv2::Streaming().BaseUrl(relay_url))
        .Synchronizer(FDv2::Streaming())
        .Synchronizer(FDv2::Polling()));

auto config = config_builder.Build();
if (!config) {
    /* an error occurred, config is not valid */
}
```
