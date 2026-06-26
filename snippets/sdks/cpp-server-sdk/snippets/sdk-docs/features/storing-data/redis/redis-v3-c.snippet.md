---
id: cpp-server-sdk/sdk-docs/features/storing-data/redis/redis-v3-c
sdk: cpp-server-sdk
kind: reference
lang: c
description: Redis source configuration example for C++ (server-side) SDK v3.0 (C binding).
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only-redis

---

```c
// Stack allocate the result struct, which will hold the result pointer or
// an error message.
struct LDServerLazyLoadRedisResult result;

// Create the Redis source, passing in arguments for the URI, prefix, and
// pointer to the result.
if (!LDServerLazyLoadRedisSource_New("redis://localhost:6379", "my-key-prefix",
&result)) {
    // Redis config is invalid, cannot proceed.
    // Error message is stored in result.error_message.
}

// Create a builder for the Lazy Load data system.
LDServerLazyLoadBuilder lazy_builder = LDServerLazyLoadBuilder_New();

// Pass the Redis source pointer into it.
LDServerLazyLoadBuilder_SourcePtr(lazy_builder,
                                  (LDServerLazyLoadSourcePtr)result.source);

// Create a standard server-side SDK configuration builder.
LDServerConfigBuilder cfg_builder = LDServerConfigBuilder_New("YOUR_SDK_KEY");

// Tell the SDK config builder to use the Lazy Load system that was just
// configured.
LDServerConfigBuilder_DataSystem_LazyLoad(cfg_builder, lazy_builder);

LDServerConfig config;
LDStatus status = LDServerConfigBuilder_Build(cfg_builder, &config);

if (!LDStatus_Ok(status)) {
    // an error occurred, config is not valid
}
```
