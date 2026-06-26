---
id: cpp-server-sdk/sdk-docs/features/relay-proxy-config/daemon-mode/daemon-mode-cpp-native-v3-0
sdk: cpp-server-sdk
kind: reference
lang: cpp
description: Daemon mode (lazy load) configuration example for C++ (server-side).
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only
---

```cpp
using LazyLoad = server_side::config::builders::LazyLoadBuilder;

auto config_builder = server_side::ConfigBuilder(sdk_key);

auto some_source = YourDatabaseIntegration();

config_builder.DataSystem().Method(
    LazyLoad().Source(some_source)
);

auto config = config_builder.Build();
if (!config) {
    /* an error occurred, config is not valid */
}
```
