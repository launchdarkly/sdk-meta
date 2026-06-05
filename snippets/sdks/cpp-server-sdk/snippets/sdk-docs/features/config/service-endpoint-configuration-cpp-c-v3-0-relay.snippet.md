---
id: cpp-server-sdk/sdk-docs/features/config/service-endpoint-configuration-cpp-c-v3-0-relay
sdk: cpp-server-sdk
kind: reference
lang: c
description: Service endpoint configuration example for C++ (server-side).
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only

---

```c
LDServerConfigBuilder builder = LDServerConfigBuilder_New("YOUR_SDK_KEY");

LDServerConfigBuilder_ServiceEndpoints_RelayProxyBaseURL(builder, "https://your-relay-proxy.com:8030");

LDServerConfig config;
LDStatus status = LDServerConfigBuilder_Build(builder, &config);
if (!LDStatus_Ok(status)) {
    /* an error occurred, config is not valid */
}
```
