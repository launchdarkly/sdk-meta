---
id: cpp-server-sdk/sdk-docs/features/config/index-cpp-c-v3-0
sdk: cpp-server-sdk
kind: reference
lang: c
description: SDK configuration example for C++ (server-side).
---

```c
LDServerConfigBuilder builder = LDServerConfigBuilder_New("YOUR_SDK_KEY");
LDServerConfigBuilder_Events_Capacity(builder, 1000);
LDServerConfigBuilder_Events_FlushIntervalMs(builder, 30 * 1000);
LDServerConfig config;

LDStatus status = LDServerConfigBuilder_Build(builder, &config);

if (!LDStatus_Ok(status)) {
    /* an error occurred, config is not valid */
}
```
