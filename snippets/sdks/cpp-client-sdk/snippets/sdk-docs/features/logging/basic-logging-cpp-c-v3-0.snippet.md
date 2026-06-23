---
id: cpp-client-sdk/sdk-docs/features/logging/basic-logging-cpp-c-v3-0
sdk: cpp-client-sdk
kind: reference
lang: c
description: Default-logger configuration example for C++ (client-side) SDK v3.0 (C binding).
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only

---

```c
LDClientConfigBuilder config_builder = LDClientConfigBuilder_New("example-mobile-key");

LDLoggingBasicBuilder basic_logging = LDLoggingBasicBuilder_New();
LDLoggingBasicBuilder_Tag(basic_logging, "ArbitraryLogTag");
LDLoggingBasicBuilder_Level(basic_logging, LD_LOG_WARN);

LDClientConfigBuilder_Logging_Basic(config_builder, basic_logging);

LDClientConfig config;
LDStatus status = LDClientConfigBuilder_Build(config_builder, &config);
```
