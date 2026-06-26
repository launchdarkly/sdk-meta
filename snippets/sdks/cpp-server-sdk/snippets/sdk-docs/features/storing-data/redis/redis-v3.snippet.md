---
id: cpp-server-sdk/sdk-docs/features/storing-data/redis/redis-v3
sdk: cpp-server-sdk
kind: reference
lang: cpp
description: Redis source configuration example for C++ (server-side) SDK v3.0 (native).
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only-redis

---

```cpp
// Make sure to include the redis source's header.
#include <launchdarkly/server_side/integrations/redis/redis_source.hpp>

using namespace launchdarkly::server_side;

using LazyLoad = config::builders::LazyLoadBuilder;

ConfigBuilder config_builder(sdk_key);

auto redis_source = integrations::RedisDataSource::Create("redis://localhost:6379", "my-key-prefix");

if (!redis_source) {
    /* redis config is invalid, cannot proceed */
}

config_builder.DataSystem().Method(
    LazyLoad().Source(std::move(*redis_source)).CacheRefresh(std::chrono::seconds(15))
);

auto config = config_builder.Build();
if (!config) {
    /* an error occurred, config is not valid */
}
```
