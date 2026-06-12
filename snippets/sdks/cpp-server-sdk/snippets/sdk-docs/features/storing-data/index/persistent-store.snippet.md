---
id: cpp-server-sdk/sdk-docs/features/storing-data/index/persistent-store
sdk: cpp-server-sdk
kind: reference
lang: cpp
description: Persistent feature store configuration example for C++ (server-side) SDK v3.0 (native).
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
