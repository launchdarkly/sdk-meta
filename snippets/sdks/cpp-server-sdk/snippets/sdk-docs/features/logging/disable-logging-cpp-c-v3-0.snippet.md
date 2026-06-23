---
id: cpp-server-sdk/sdk-docs/features/logging/disable-logging-cpp-c-v3-0
sdk: cpp-server-sdk
kind: reference
lang: c
description: Disable-logging example for C++ (server-side) SDK v3.0 (C binding).
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only

---

```c
LDServerConfigBuilder builder = LDServerConfigBuilder_New("YOUR_SDK_KEY");

LDServerConfigBuilder_Logging_Disable(builder);

LDServerConfig config;
LDStatus status = LDServerConfigBuilder_Build(builder, &config);
```
