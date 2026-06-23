---
id: cpp-server-sdk/sdk-docs/features/logging/custom-logger-backend-cpp-native-v3-0
sdk: cpp-server-sdk
kind: reference
lang: cpp
description: Custom log backend implementation example for C++ (server-side) SDK v3.0 (native).
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only-toplevel

---

```cpp
#include <launchdarkly/logging/log_level.hpp>
#include <launchdarkly/logging/log_backend.hpp>

using namespace launchdarkly;

class CustomLogger : public ILogBackend {
   public:
    /* Should return true if the specified level is enabled; in this example, return true to log all messages. */
    bool Enabled(LogLevel level) noexcept override { return true; }

    /* Forwards to stdout as an example, printing the log tag along with the message. */
    void Write(LogLevel level, std::string message) noexcept override {
        std::cout << GetLogLevelName(level, "unknown") << ": " << message << std::endl;
}
};
```
