---
id: cpp-client-sdk/sdk-docs/features/config/service-endpoint-configuration-cpp-native-v3-0-federal
sdk: cpp-client-sdk
kind: reference
lang: cpp
description: Service endpoint configuration example for C++ (client-side).
---

```cpp
auto config_builder = client_side::ConfigBuilder("example-mobile-key");
config_builder.ServiceEndpoints()
    .StreamingBaseUrl("https://clientstream.launchdarkly.us")
    .PollingBaseUrl("https://clientsdk.launchdarkly.us")
    .EventsBaseUrl("https://events.launchdarkly.us")
auto config = config_builder.Build();
```
