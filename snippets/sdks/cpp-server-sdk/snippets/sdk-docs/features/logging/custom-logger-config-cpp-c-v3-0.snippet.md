---
id: cpp-server-sdk/sdk-docs/features/logging/custom-logger-config-cpp-c-v3-0
sdk: cpp-server-sdk
kind: reference
lang: c
description: Custom logger installation example for C++ (server-side) SDK v3.0 (C binding).
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only

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

LDServerConfigBuilder builder = LDServerConfigBuilder_New("YOUR_SDK_KEY");
LDServerConfigBuilder_Logging_Custom(builder, custom_logging);

LDServerConfig config;
LDStatus status = LDServerConfigBuilder_Build(builder, &config);
```
