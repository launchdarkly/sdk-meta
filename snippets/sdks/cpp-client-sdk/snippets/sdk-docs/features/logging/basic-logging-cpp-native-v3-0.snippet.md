---
id: cpp-client-sdk/sdk-docs/features/logging/basic-logging-cpp-native-v3-0
sdk: cpp-client-sdk
kind: reference
lang: cpp
description: Default-logger configuration example for C++ (client-side) SDK v3.0 (native).
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only

---

```cpp
auto config_builder = client_side::ConfigBuilder("example-mobile-key");
config_builder.Logging()
  .Logging(LoggingBuilder::BasicLogging().Tag("ArbitraryLogTag").Level(LogLevel::kWarn));
auto config = config_builder.Build();
```
