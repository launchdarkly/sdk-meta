---
id: cpp-server-sdk/sdk-docs/features/logging/custom-logger-config-cpp-native-v3-0
sdk: cpp-server-sdk
kind: reference
lang: cpp
description: Custom logger installation example for C++ (server-side) SDK v3.0 (native).
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only

---

```cpp
// Make sure the <memory> header is included for std::make_shared
#include <memory>

auto config_builder = server_side::ConfigBuilder("YOUR_SDK_KEY");

config_builder.Logging().Logging(LoggingBuilder::CustomLogging().Backend(
  std::make_shared<CustomLogger>()));

auto config = config_builder.Build();
```
