---
id: cpp-server-sdk/scaffolds/cpp-syntax-only-dynamodb
sdk: cpp-server-sdk
kind: scaffold
lang: cpp
file: main.cpp
description: |
  Parse-only validator for C++ server SDK doc fragments that reference
  the DynamoDB integration -- both the Lazy Load persistent store
  (`DynamoDBDataSource`) and the Big Segments store
  (`DynamoDBBigSegmentStore`).

  Same shape as `cpp-syntax-only-redis`, with one difference in how the
  build is wired. The DynamoDB source library links the AWS C++ SDK,
  which is fetched from source and is too heavy to build in CI. But its
  public headers are AWS-clean (they pull in no AWS headers), and the
  wrappee below is a never-instantiated template, so no DynamoDB or AWS
  symbols are ever ODR-used. The cpp-server Dockerfile therefore puts
  the DynamoDB source library's `include/` directory on the compile
  path without enabling `LD_BUILD_DYNAMODB_SUPPORT` -- the fragment's
  API calls type-check against the real headers with no AWS SDK build
  and no link. No `CPP_*` env flag is set: these fragments compile in
  the default project alongside the other `cpp-syntax-only` snippets.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: cpp-server
  entrypoint: main.cpp
---

```cpp
#include <chrono>
#include <cstdio>
#include <future>
#include <iostream>
#include <memory>
#include <optional>
#include <string>
// Native C++ headers.
#include <launchdarkly/server_side/client.hpp>
#include <launchdarkly/server_side/config/config_builder.hpp>
#include <launchdarkly/context_builder.hpp>
#include <launchdarkly/value.hpp>
// C-binding headers — doc fragments mix C-binding and native styles.
#include <launchdarkly/server_side/bindings/c/sdk.h>
#include <launchdarkly/server_side/bindings/c/config/builder.h>
#include <launchdarkly/bindings/c/context.h>
#include <launchdarkly/bindings/c/context_builder.h>
// DynamoDB source + Big Segments store headers (native + C binding),
// plus the core Big Segments builder the store is handed to. Bodies
// that #include these themselves hit the include guard and stay
// verbatim.
#include <launchdarkly/server_side/integrations/dynamodb/dynamodb_source.hpp>
#include <launchdarkly/server_side/bindings/c/integrations/dynamodb/dynamodb_source.h>
#include <launchdarkly/server_side/integrations/dynamodb/dynamodb_big_segment_store.hpp>
#include <launchdarkly/server_side/bindings/c/integrations/dynamodb/dynamodb_big_segment_store.h>
#include <launchdarkly/server_side/config/builders/big_segments_builder.hpp>
#include <launchdarkly/server_side/bindings/c/config/big_segments_builder/big_segments_builder.h>

// Wrappee is a never-instantiated template — body is parsed but
// most type-checks are deferred to instantiation (which never
// happens). The body lives in a nested block so it can re-declare
// names the scaffold stubs. As in cpp-syntax-only-redis, no
// namespaces are lifted here: the native DynamoDB fragment carries
// its own `using namespace launchdarkly::server_side;`. Init-shaped
// fragments pass an `sdk_key` the docs assume already exists.
template <int = 0>
void _wrappee() {
    char const* sdk_key = "";
    (void)sdk_key;
    {
{{ body }}
    }
}

int main() {
    std::cout << "feature flag evaluates to true" << std::endl;
    return 0;
}
```
