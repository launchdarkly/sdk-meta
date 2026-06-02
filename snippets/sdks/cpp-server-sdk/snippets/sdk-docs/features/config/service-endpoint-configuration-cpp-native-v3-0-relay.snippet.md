---
id: cpp-server-sdk/sdk-docs/features/config/service-endpoint-configuration-cpp-native-v3-0-relay
sdk: cpp-server-sdk
kind: reference
lang: cpp
description: Service endpoint configuration example for C++ (server-side).
---

```cpp
auto config_builder = server_side::ConfigBuilder("YOUR_SDK_KEY");
config_builder.ServiceEndpoints()
    .RelayProxyBaseUrl("https://your-relay-proxy.com:8030");
auto config = config_builder.Build();
if (!config) {
    /* an error occurred, config is not valid */
}
```
