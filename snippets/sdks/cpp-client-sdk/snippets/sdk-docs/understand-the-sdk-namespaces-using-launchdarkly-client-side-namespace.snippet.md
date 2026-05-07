---
id: cpp-client-sdk/sdk-docs/understand-the-sdk-namespaces-using-launchdarkly-client-side-namespace
sdk: cpp-client-sdk
kind: reference
lang: cpp
description: "Using launchdarkly::client_side namespace in section \"Understand the SDK namespaces\""
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only
---

```cpp
using namespace launchdarkly::client_side;
auto config_builder = ConfigBuilder("example-mobile-key");
auto config = config_builder.Build();
```
