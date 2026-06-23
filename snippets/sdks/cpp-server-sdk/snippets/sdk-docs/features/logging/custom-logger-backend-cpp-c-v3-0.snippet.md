---
id: cpp-server-sdk/sdk-docs/features/logging/custom-logger-backend-cpp-c-v3-0
sdk: cpp-server-sdk
kind: reference
lang: c
description: Custom log backend callback definitions example for C++ (server-side) SDK v3.0 (C binding).
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only-toplevel

---

```c

/* Should return true if the specified level is enabled; in this example, return true to log all messages. */
static bool enabled(enum LDLogLevel level, void *user_data) {
    return true;
}

/* Forwards to stdout as an example, printing the log tag along with the message. */
static void write(enum LDLogLevel level, const char *msg, void *user_data) {
    printf("%d: %s\n", level, msg);
}
```
