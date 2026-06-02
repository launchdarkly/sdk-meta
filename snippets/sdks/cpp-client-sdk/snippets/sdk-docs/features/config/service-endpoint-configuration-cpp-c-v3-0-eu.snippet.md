---
id: cpp-client-sdk/sdk-docs/features/config/service-endpoint-configuration-cpp-c-v3-0-eu
sdk: cpp-client-sdk
kind: reference
lang: c
description: Service endpoint configuration example for C++ (client-side).
---

```c
LDClientConfigBuilder builder = LDClientConfigBuilder_New("example-mobile-key");
LDClientConfigBuilder_ServiceEndpoints_StreamingBaseURL(builder, "https://clientstream.eu.launchdarkly.com")
LDClientConfigBuilder_ServiceEndpoints_PollingBaseURL(builder, "https://clientsdk.eu.launchdarkly.com")
LDClientConfigBuilder_ServiceEndpoints_EventsBaseURL(builder, "https://events.eu.launchdarkly.com")
LDClientConfig config;

LDStatus status = LDClientConfigBuilder_Build(builder, &config);

if (!LDStatus_Ok(status)) {
    /* an error occurred, config is not valid */
}
```
