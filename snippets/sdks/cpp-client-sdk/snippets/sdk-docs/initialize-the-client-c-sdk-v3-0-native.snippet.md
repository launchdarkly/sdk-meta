---
id: cpp-client-sdk/sdk-docs/initialize-the-client-c-sdk-v3-0-native
sdk: cpp-client-sdk
kind: reference
lang: cpp
description: "C++ SDK v3.0 (native) in section \"Initialize the client\""
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only
---

```cpp

auto config_builder = client_side::ConfigBuilder("example-mobile-key");
auto config = config_builder.Build();
if (!config) {
   /* an error occurred, config is not valid */
}
auto context = ContextBuilder().Kind("user", "example-user-key").Build();
```
