---
id: cpp-server-sdk/sdk-docs/migration-2-to-3-understanding-changes-to-storing-data-3-0-syntax-use-lazy-loading
sdk: cpp-server-sdk
kind: reference
lang: cpp
description: "3.0 syntax, use lazy loading in section \"Understanding changes to storing data\""
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
