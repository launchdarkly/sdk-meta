---
id: cpp-client-sdk/sdk-docs/features/config/index-cpp-c-v3-0
sdk: cpp-client-sdk
kind: reference
lang: c
description: SDK configuration example for C++ (client-side).
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only

---

```c
LDClientConfigBuilder builder = LDClientConfigBuilder_New("example-mobile-key");
LDClientConfigBuilder_Events_Capacity(builder, 1000);
LDClientConfigBuilder_Events_FlushIntervalMs(builder, 30 * 1000);
LDClientConfig config;

LDStatus status = LDClientConfigBuilder_Build(builder, &config);

if (!LDStatus_Ok(status)) {
    /* an error occurred, config is not valid */
}
```
