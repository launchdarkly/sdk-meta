---
id: cpp-client-sdk/sdk-docs/features/config/service-endpoint-configuration-cpp-native-v3-0-relay
sdk: cpp-client-sdk
kind: reference
lang: cpp
description: Service endpoint configuration example for C++ (client-side).
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only

---

```cpp
auto config_builder = client_side::ConfigBuilder("example-mobile-key");
config_builder.ServiceEndpoints()
    .StreamingBaseUrl("https://your-relay-proxy.com:8030")
    .PollingBaseUrl("https://your-relay-proxy.com:8030")
    .EventsBaseUrl("https://your-relay-proxy.com:8030");
auto config = config_builder.Build();
```
