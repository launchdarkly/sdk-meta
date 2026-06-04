---
id: cpp-client-sdk/sdk-docs/features/config/service-endpoint-configuration-cpp-c-v3-0-federal
sdk: cpp-client-sdk
kind: reference
lang: c
description: Service endpoint configuration example for C++ (client-side).
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only

---

```c
LDClientConfigBuilder builder = LDClientConfigBuilder_New("example-mobile-key");
LDClientConfigBuilder_ServiceEndpoints_StreamingBaseURL(builder, "https://clientstream.launchdarkly.us");
LDClientConfigBuilder_ServiceEndpoints_PollingBaseURL(builder, "https://clientsdk.launchdarkly.us");
LDClientConfigBuilder_ServiceEndpoints_EventsBaseURL(builder, "https://events.launchdarkly.us");
LDClientConfig config;

LDStatus status = LDClientConfigBuilder_Build(builder, &config);

if (!LDStatus_Ok(status)) {
    /* an error occurred, config is not valid */
}
```
