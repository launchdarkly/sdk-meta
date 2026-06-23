---
id: cpp-server-sdk/sdk-docs/features/logging/basic-logging-cpp-native-v3-0
sdk: cpp-server-sdk
kind: reference
lang: cpp
description: Default-logger configuration example for C++ (server-side) SDK v3.0 (native).
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only

---

```cpp
auto config_builder = server_side::ConfigBuilder("YOUR_SDK_KEY");

using LoggingBuilder = server_side::config::builders::LoggingBuilder;
config_builder.Logging().Logging(
  LoggingBuilder::BasicLogging().Tag("ArbitraryLogTag").Level(LogLevel::kWarn)
);
```
