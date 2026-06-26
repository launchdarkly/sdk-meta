---
id: cpp-server-sdk/sdk-docs/features/logging/custom-logger-backend-c-sdk-v2
sdk: cpp-server-sdk
kind: reference
lang: c
description: Custom logger function definition example for the C server SDK v2.x.
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only-v2-c-toplevel

---

```c
static void
myCustomLogger(const LDLogLevel level, const char *const text)
{
    printf("[%s] %s\n", LDLogLevelToString(level), text);
}

```
