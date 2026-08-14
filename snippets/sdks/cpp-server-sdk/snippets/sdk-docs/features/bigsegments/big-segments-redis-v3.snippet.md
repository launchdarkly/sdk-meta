---
id: cpp-server-sdk/sdk-docs/features/bigsegments/big-segments-redis-v3
sdk: cpp-server-sdk
kind: reference
lang: cpp
description: Big Segments Redis store configuration example for C++ (server-side) SDK v3.0 (native).
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only-redis

---

```cpp
// Make sure to include the Redis Big Segments store's header.
#include <launchdarkly/server_side/integrations/redis/redis_big_segment_store.hpp>

using namespace launchdarkly::server_side;

ConfigBuilder config_builder(sdk_key);

auto redis_store = integrations::RedisBigSegmentStore::Create("redis://localhost:6379", "my-key-prefix");

if (!redis_store) {
    /* redis config is invalid, cannot proceed */
}

config_builder.BigSegments(
    config::builders::BigSegmentsBuilder(std::move(*redis_store))
        .ContextCacheSize(2000)
        .ContextCacheTime(std::chrono::seconds(30))
);

auto config = config_builder.Build();
if (!config) {
    /* an error occurred, config is not valid */
}
```
