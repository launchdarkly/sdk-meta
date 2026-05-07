---
id: cpp-client-sdk/sdk-docs/understand-the-sdk-namespaces-using-launchdarkly-namespace
sdk: cpp-client-sdk
kind: reference
lang: cpp
description: "Using launchdarkly namespace in section \"Understand the SDK namespaces\""
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only
---

```cpp
using namespace launchdarkly; // omitted in examples; assumed to be present
auto config_builder = client_side::ConfigBuilder("example-mobile-key");
auto config = config_builder.Build();
```
