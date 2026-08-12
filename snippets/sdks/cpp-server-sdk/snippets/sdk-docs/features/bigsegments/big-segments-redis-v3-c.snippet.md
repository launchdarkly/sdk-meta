---
id: cpp-server-sdk/sdk-docs/features/bigsegments/big-segments-redis-v3-c
sdk: cpp-server-sdk
kind: reference
lang: c
description: Big Segments Redis store configuration example for C++ (server-side) SDK v3.0 (C binding).
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only-redis

---

```c
// Stack allocate the result struct, which will hold the store pointer or
// an error message.
struct LDServerBigSegmentsRedisResult result;

// Create the Redis Big Segments store, passing in arguments for the URI,
// prefix, and pointer to the result.
if (!LDServerBigSegmentsRedisStore_New("redis://localhost:6379", "my-key-prefix",
                                       &result)) {
    // Redis config is invalid, cannot proceed.
    // Error message is stored in result.error_message.
}

// Create a Big Segments builder from the store pointer.
LDServerBigSegmentsBuilder bs_builder =
    LDServerBigSegmentsBuilder_New((LDServerBigSegmentStorePtr)result.store);

LDServerBigSegmentsBuilder_ContextCacheSize(bs_builder, 2000);
LDServerBigSegmentsBuilder_ContextCacheTimeMs(bs_builder, 30000);

// Create a standard server-side SDK configuration builder.
LDServerConfigBuilder cfg_builder = LDServerConfigBuilder_New("YOUR_SDK_KEY");

// Tell the SDK config builder to use the Big Segments store that was just
// configured.
LDServerConfigBuilder_BigSegments(cfg_builder, bs_builder);

LDServerConfig config;
LDStatus status = LDServerConfigBuilder_Build(cfg_builder, &config);

if (!LDStatus_Ok(status)) {
    // an error occurred, config is not valid
}
```
