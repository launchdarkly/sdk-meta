---
id: cpp-server-sdk/sdk-docs/features/config/index-c-sdk-v2
sdk: cpp-server-sdk
kind: reference
lang: c
description: SDK configuration example for C++ (server-side).
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only-v2-c

---

```c
struct LDConfig *config = LDConfigNew("YOUR_SDK_KEY");
LDConfigSetEventsCapacity(config, 1000);
LDConfigSetEventsFlushInterval(config, 30000);
```
