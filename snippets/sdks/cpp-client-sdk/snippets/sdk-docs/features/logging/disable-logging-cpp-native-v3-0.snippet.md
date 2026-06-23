---
id: cpp-client-sdk/sdk-docs/features/logging/disable-logging-cpp-native-v3-0
sdk: cpp-client-sdk
kind: reference
lang: cpp
description: Disable-logging example for C++ (client-side) SDK v3.0 (native).
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only

---

```cpp
auto config_builder = client_side::ConfigBuilder("example-mobile-key");
config_builder.Logging().Logging(LoggingBuilder::NoLogging());

auto config = config_builder.Build();
```
