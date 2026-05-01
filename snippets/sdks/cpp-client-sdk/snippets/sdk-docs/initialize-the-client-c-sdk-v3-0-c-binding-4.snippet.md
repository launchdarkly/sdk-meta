---
id: cpp-client-sdk/sdk-docs/initialize-the-client-c-sdk-v3-0-c-binding-4
sdk: cpp-client-sdk
kind: reference
lang: c
description: "C++ SDK v3.0 (C binding) in section \"Initialize the client\""
---

```c
LDClientSDK client = LDClientSDK_New(config, context);

bool initialized_successfully;
if (LDClientSDK_Start(client, maxwait, &initialized_successfully)) {
  /* The client's attempt to initialize succeeded or failed in the specified amount of time. */
  if (initialized_successfully) {
    /* Initialization succeeded. */
  else {
    /* Initialization failed. */
  }
} else {
   /* The specified timeout was reached, but the client is still initializing. */
}
```
