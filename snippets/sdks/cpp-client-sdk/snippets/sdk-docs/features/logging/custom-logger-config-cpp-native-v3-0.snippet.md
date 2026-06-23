---
id: cpp-client-sdk/sdk-docs/features/logging/custom-logger-config-cpp-native-v3-0
sdk: cpp-client-sdk
kind: reference
lang: cpp
description: Custom logger installation example for C++ (client-side) SDK v3.0 (native).
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only

---

```cpp
// Make sure the <memory> header is included for std::make_shared
#include <memory>

auto config_builder = client_side::ConfigBuilder("example-mobile-key");
config_builder.Logging()
  .Logging(LoggingBuilder::CustomLogging().Backend(std::make_shared<CustomLogger>()));
auto config = config_builder.Build();
```
