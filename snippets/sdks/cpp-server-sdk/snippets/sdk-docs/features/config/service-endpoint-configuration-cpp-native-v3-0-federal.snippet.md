---
id: cpp-server-sdk/sdk-docs/features/config/service-endpoint-configuration-cpp-native-v3-0-federal
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
    .StreamingBaseUrl("https://stream.launchdarkly.us")
    .PollingBaseUrl("https://sdk.launchdarkly.us")
    .EventsBaseUrl("https://events.launchdarkly.us");
auto config = config_builder.Build();
if (!config) {
    /* an error occurred, config is not valid */
}
```
