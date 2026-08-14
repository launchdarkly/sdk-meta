---
id: cpp-server-sdk/sdk-docs/features/storing-data/dynamodb/dynamodb-v3
sdk: cpp-server-sdk
kind: reference
lang: cpp
description: DynamoDB source configuration example for C++ (server-side) SDK v3.0 (native).
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only-dynamodb

---

```cpp
// Make sure to include the DynamoDB source's header.
#include <launchdarkly/server_side/integrations/dynamodb/dynamodb_source.hpp>

using namespace launchdarkly::server_side;

using LazyLoad = config::builders::LazyLoadBuilder;

ConfigBuilder config_builder(sdk_key);

auto dynamodb_source = integrations::DynamoDBDataSource::Create("my-table", "my-key-prefix");

if (!dynamodb_source) {
    /* dynamodb config is invalid, cannot proceed */
}

config_builder.DataSystem().Method(
    LazyLoad().Source(std::move(*dynamodb_source)).CacheRefresh(std::chrono::seconds(15))
);

auto config = config_builder.Build();
if (!config) {
    /* an error occurred, config is not valid */
}
```
