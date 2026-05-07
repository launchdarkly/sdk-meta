---
id: cpp-server-sdk/sdk-docs/c-c---about-the-sdk-namespaces-using-launchdarkly-namespace
sdk: cpp-server-sdk
kind: reference
lang: cpp
description: "Using launchdarkly namespace in section \"About the SDK namespaces\""
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only
---

```cpp
using namespace launchdarkly; // omitted in examples; assumed to be present
auto config_builder = server_side::ConfigBuilder("YOUR_SDK_KEY");
auto config = config_builder.Build();
```
