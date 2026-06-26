---
id: cpp-server-sdk/sdk-docs/features/relay-proxy-config/proxy-mode/proxy-mode-cpp-native-v3-0
sdk: cpp-server-sdk
kind: reference
lang: cpp
description: Proxy mode configuration example for C++ (server-side).
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only
---

```cpp
auto config_builder = server_side::ConfigBuilder("YOUR_SDK_KEY");
config_builder.ServiceEndpoints().RelayProxyBaseURL("https://your-relay-proxy.com:8030");
auto config = config_builder.Build();
if (!config) {
   /* an error occurred, config is not valid */
}
```
