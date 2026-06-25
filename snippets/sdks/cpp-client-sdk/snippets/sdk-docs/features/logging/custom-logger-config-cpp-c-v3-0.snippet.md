---
id: cpp-client-sdk/sdk-docs/features/logging/custom-logger-config-cpp-c-v3-0
sdk: cpp-client-sdk
kind: reference
lang: c
description: Custom logger installation example for C++ (client-side) SDK v3.0 (C binding).
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only

---

```c
struct LDLogBackend backend;
LDLogBackend_Init(&backend);

backend.Write = write;
backend.Enabled = enabled;

/* You may optionally assign the UserData pointer, which will be passed into Write and Enabled. */
/* backend.UserData = &some_struct; */

LDLoggingCustomBuilder custom_logging = LDLoggingCustomBuilder_New();
LDLoggingCustomBuilder_Backend(custom_logging, backend);

LDClientConfigBuilder builder = LDClientConfigBuilder_New("example-mobile-key");
LDClientConfigBuilder_Logging_Custom(builder, custom_logging);

LDClientConfig config;
LDStatus status = LDClientConfigBuilder_Build(builder, &config);
```
