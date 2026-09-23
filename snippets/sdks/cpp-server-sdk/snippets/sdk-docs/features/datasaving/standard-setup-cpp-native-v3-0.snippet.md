---
id: cpp-server-sdk/sdk-docs/features/datasaving/standard-setup-cpp-native-v3-0
sdk: cpp-server-sdk
kind: reference
lang: cpp
description: Data saving mode standard setup for C++ (server-side).
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only
---

```cpp
using FDv2 = server_side::config::builders::DataSystemBuilder::FDv2;

auto config_builder = server_side::ConfigBuilder("YOUR_SDK_KEY");
config_builder.DataSystem().Method(FDv2::Default());

auto config = config_builder.Build();
if (!config) {
    /* an error occurred, config is not valid */
}
```
