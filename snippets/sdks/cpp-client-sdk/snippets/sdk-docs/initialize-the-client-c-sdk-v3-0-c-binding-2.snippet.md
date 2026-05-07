---
id: cpp-client-sdk/sdk-docs/initialize-the-client-c-sdk-v3-0-c-binding-2
sdk: cpp-client-sdk
kind: reference
lang: c
description: "C++ SDK v3.0 (C binding) in section \"Initialize the client\""
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only
---

```c
LDClientSDK client = LDClientSDK_New(config, context);

unsigned int maxwait = 10 * 1000; /* 10 seconds */
LDClientSDK_Start(client, maxwait, NULL);
```
