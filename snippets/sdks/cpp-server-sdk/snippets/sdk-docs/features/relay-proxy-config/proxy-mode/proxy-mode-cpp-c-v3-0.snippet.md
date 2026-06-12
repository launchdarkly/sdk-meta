---
id: cpp-server-sdk/sdk-docs/features/relay-proxy-config/proxy-mode/proxy-mode-cpp-c-v3-0
sdk: cpp-server-sdk
kind: reference
lang: c
description: Proxy mode configuration example for C++ (server-side, C binding).
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
