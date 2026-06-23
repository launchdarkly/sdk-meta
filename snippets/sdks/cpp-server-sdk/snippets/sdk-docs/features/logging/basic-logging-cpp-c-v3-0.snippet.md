---
id: cpp-server-sdk/sdk-docs/features/logging/basic-logging-cpp-c-v3-0
sdk: cpp-server-sdk
kind: reference
lang: c
description: Default-logger configuration example for C++ (server-side) SDK v3.0 (C binding).
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only

---

```c
LDServerConfigBuilder config_builder = LDServerConfigBuilder_New("YOUR_SDK_KEY");

LDLoggingBasicBuilder basic_logging = LDLoggingBasicBuilder_New();
LDLoggingBasicBuilder_Tag(basic_logging, "ArbitraryLogTag");
LDLoggingBasicBuilder_Level(basic_logging, LD_LOG_WARN);

LDServerConfigBuilder_Logging_Basic(config_builder, basic_logging);

LDServerConfig config;
LDStatus status = LDServerConfigBuilder_Build(config_builder, &config);
```
