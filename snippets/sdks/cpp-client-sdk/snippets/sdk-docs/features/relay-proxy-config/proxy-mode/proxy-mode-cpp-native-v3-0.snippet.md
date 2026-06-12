---
id: cpp-client-sdk/sdk-docs/features/relay-proxy-config/proxy-mode/proxy-mode-cpp-native-v3-0
sdk: cpp-client-sdk
kind: reference
lang: cpp
description: Proxy mode configuration example for C++ (client-side).
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only
---

```cpp
auto config_builder = client_side::ConfigBuilder("example-mobile-key");
config_builder.ServiceEndpoints().RelayProxyBaseURL("https://your-relay-proxy.com:8030");
auto config = config_builder.Build();
```
