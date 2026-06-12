---
id: cpp-client-sdk/sdk-docs/features/logging/disable-logging-cpp-c-v3-0
sdk: cpp-client-sdk
kind: reference
lang: c
description: Disable-logging example for C++ (client-side) SDK v3.0 (C binding).
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only

---

```c
LDClientConfigBuilder builder = LDClientConfigBuilder_New("example-mobile-key");

LDClientConfigBuilder_Logging_Disable(builder);

LDClientConfig config;
LDStatus status = LDClientConfigBuilder_Build(builder, &config);
```
