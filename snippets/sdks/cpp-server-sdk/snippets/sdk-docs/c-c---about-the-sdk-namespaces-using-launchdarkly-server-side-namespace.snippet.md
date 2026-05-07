---
id: cpp-server-sdk/sdk-docs/c-c---about-the-sdk-namespaces-using-launchdarkly-server-side-namespace
sdk: cpp-server-sdk
kind: reference
lang: cpp
description: "Using launchdarkly::server_side namespace in section \"About the SDK namespaces\""
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only
---

```cpp
using namespace launchdarkly::server_side;
auto config_builder = ConfigBuilder("YOUR_SDK_KEY");
auto config = config_builder.Build();
```
