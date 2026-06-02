---
id: cpp-server-sdk/sdk-docs/features/config/app-config-cpp-c-v3-0
sdk: cpp-server-sdk
kind: reference
lang: c
description: Application metadata configuration example for C++ (server-side).
---

```c
LDServerConfigBuilder builder = LDServerConfigBuilder_New("YOUR_SDK_KEY");
LDServerConfigBuilder_AppInfo_Identifier(builder, "authentication-service");
LDServerConfigBuilder_AppInfo_Version(builder, "1.0.0");
LDServerConfig config;

LDStatus status = LDServerConfigBuilder_Build(builder, &config);

if (!LDStatus_Ok(status)) {
    /* an error occurred, config is not valid */
}
```
