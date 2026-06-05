---
id: cpp-client-sdk/sdk-docs/features/config/index-c-sdk-v2
sdk: cpp-client-sdk
kind: reference
lang: c
description: SDK configuration example for C++ (client-side).
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only-v2-c

---

```c
struct LDConfig *config = LDConfigNew("example-mobile-key");
LDConfigSetEventsCapacity(config, 1000);
LDConfigSetEventsFlushIntervalMillis(config, 30000);
```
