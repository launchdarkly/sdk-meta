---
id: cpp-client-sdk/sdk-docs/features/config/app-config-cpp-c-v3-0
sdk: cpp-client-sdk
kind: reference
lang: c
description: Application metadata configuration example for C++ (client-side).
---

```c
LDClientConfigBuilder builder = LDClientConfigBuilder_New("example-mobile-key");
LDClientConfigBuilder_AppInfo_Identifier(builder, "authentication-service");
LDClientConfigBuilder_AppInfo_Version(builder, "1.0.0");
LDClientConfig config;

LDStatus status = LDClientConfigBuilder_Build(builder, &config);

if (!LDStatus_Ok(status)) {
    /* an error occurred, config is not valid */
}
```
