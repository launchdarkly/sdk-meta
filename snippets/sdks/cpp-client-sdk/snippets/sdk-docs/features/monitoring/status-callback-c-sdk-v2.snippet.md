---
id: cpp-client-sdk/sdk-docs/features/monitoring/status-callback-c-sdk-v2
sdk: cpp-client-sdk
kind: reference
lang: c
description: Client status callback for the C SDK v2.x (native).
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only-v2-c

---

```c
void initCallback(LDStatus status)
{
    if (status == LDStatusInitialized) {
        printf("Completed LaunchDarkly client initialization");
    }
}

LDSetClientStatusCallback(initCallback);
```
