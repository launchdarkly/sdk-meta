---
id: cpp-client-sdk/sdk-docs/initialize-the-client-c-sdk-v2-x-native-2
sdk: cpp-client-sdk
kind: reference
lang: c
description: "C SDK v2.x (native) in section \"Initialize the client\""
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only-v2-c
---

```c
unsigned int maxwait = 10 * 1000; /* 10 seconds */
struct LDClient *client = LDClientInit(config, user, maxwait);
```
