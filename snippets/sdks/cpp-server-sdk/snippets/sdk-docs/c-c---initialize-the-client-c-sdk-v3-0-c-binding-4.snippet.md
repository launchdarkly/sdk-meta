---
id: cpp-server-sdk/sdk-docs/c-c---initialize-the-client-c-sdk-v3-0-c-binding-4
sdk: cpp-server-sdk
kind: reference
lang: c
description: "C++ SDK v3.0 (C binding) in section \"Initialize the client\""
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only
---

```c
LDServerSDK client = LDServerSDK_New(config);

bool initialized_successfully;
if (LDServerSDK_Start(client, maxwait, &initialized_successfully)) {
  /* The client's attempt to initialize succeeded or failed in the specified amount of time. */
  if (initialized_successfully) {
    /* Initialization succeeded. */
  } else {
    /* Initialization failed. */
  }
} else {
   /* The specified timeout was reached, but the client is still initializing. */
}
```
