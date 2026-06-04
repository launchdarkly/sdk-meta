---
id: cpp-server-sdk/sdk-docs/features/config/service-endpoint-configuration-cpp-native-v3-0-eu
sdk: cpp-server-sdk
kind: reference
lang: cpp
description: Service endpoint configuration example for C++ (server-side).
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only

---

```cpp
auto config_builder = server_side::ConfigBuilder("YOUR_SDK_KEY");
config_builder.ServiceEndpoints()
    .StreamingBaseUrl("https://stream.eu.launchdarkly.com")
    .PollingBaseUrl("https://sdk.eu.launchdarkly.com")
    .EventsBaseUrl("https://events.eu.launchdarkly.com");
auto config = config_builder.Build();
if (!config) {
    /* an error occurred, config is not valid */
}
```
