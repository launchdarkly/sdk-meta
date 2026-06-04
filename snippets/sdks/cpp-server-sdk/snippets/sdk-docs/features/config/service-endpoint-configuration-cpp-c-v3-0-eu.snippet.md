---
id: cpp-server-sdk/sdk-docs/features/config/service-endpoint-configuration-cpp-c-v3-0-eu
sdk: cpp-server-sdk
kind: reference
lang: c
description: Service endpoint configuration example for C++ (server-side).
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only

---

```c
LDServerConfigBuilder builder = LDServerConfigBuilder_New("YOUR_SDK_KEY");
LDServerConfigBuilder_ServiceEndpoints_StreamingBaseURL(builder, "https://stream.eu.launchdarkly.com");
LDServerConfigBuilder_ServiceEndpoints_PollingBaseURL(builder, "https://sdk.eu.launchdarkly.com");
LDServerConfigBuilder_ServiceEndpoints_EventsBaseURL(builder, "https://events.eu.launchdarkly.com");
LDServerConfig config;

LDStatus status = LDServerConfigBuilder_Build(builder, &config);

if (!LDStatus_Ok(status)) {
    /* an error occurred, config is not valid */
}
```
