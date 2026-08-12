---
id: cpp-server-sdk/sdk-docs/features/bigsegments/big-segments-dynamodb-v3
sdk: cpp-server-sdk
kind: reference
lang: cpp
description: Big Segments DynamoDB store configuration example for C++ (server-side) SDK v3.0 (native).
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only-dynamodb

---

```cpp
// Make sure to include the DynamoDB Big Segments store's header.
#include <launchdarkly/server_side/integrations/dynamodb/dynamodb_big_segment_store.hpp>

using namespace launchdarkly::server_side;

ConfigBuilder config_builder(sdk_key);

auto dynamodb_store = integrations::DynamoDBBigSegmentStore::Create("my-table", "my-key-prefix");

if (!dynamodb_store) {
    /* dynamodb config is invalid, cannot proceed */
}

config_builder.BigSegments(
    config::builders::BigSegmentsBuilder(std::move(*dynamodb_store))
        .ContextCacheSize(2000)
        .ContextCacheTime(std::chrono::seconds(30))
);

auto config = config_builder.Build();
if (!config) {
    /* an error occurred, config is not valid */
}
```
