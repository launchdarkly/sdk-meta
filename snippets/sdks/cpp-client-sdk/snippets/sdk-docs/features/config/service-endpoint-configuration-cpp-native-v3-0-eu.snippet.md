---
id: cpp-client-sdk/sdk-docs/features/config/service-endpoint-configuration-cpp-native-v3-0-eu
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
    .StreamingBaseUrl("https://clientstream.eu.launchdarkly.com")
    .PollingBaseUrl("https://clientsdk.eu.launchdarkly.com")
    .EventsBaseUrl("https://events.eu.launchdarkly.com");
auto config = config_builder.Build();
```
