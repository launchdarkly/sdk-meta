---
id: cpp-server-sdk/sdk-docs/features/datasaving/standard-setup-cpp-c-v3-0
sdk: cpp-server-sdk
kind: reference
lang: c
description: Data saving mode standard setup for C++ (server-side).
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only
---

```c
LDServerConfigBuilder builder = LDServerConfigBuilder_New("YOUR_SDK_KEY");

LDServerFDv2Builder fdv2 = LDServerFDv2Builder_Default();
LDServerConfigBuilder_DataSystem_FDv2(builder, fdv2);

LDServerConfig config;
LDStatus status = LDServerConfigBuilder_Build(builder, &config);

if (!LDStatus_Ok(status)) {
    /* an error occurred, config is not valid */
}
```
