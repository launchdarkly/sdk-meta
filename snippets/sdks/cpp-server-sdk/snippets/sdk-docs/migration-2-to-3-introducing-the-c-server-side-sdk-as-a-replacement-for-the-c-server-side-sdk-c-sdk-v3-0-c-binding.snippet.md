---
id: cpp-server-sdk/sdk-docs/migration-2-to-3-introducing-the-c-server-side-sdk-as-a-replacement-for-the-c-server-side-sdk-c-sdk-v3-0-c-binding
sdk: cpp-server-sdk
kind: reference
lang: c
description: "C++ SDK v3.0 (C binding) in section \"Introducing the C++ (server-side) SDK as a replacement for the C (server-side) SDK\""
---

```c
bool initialized_successfully;
if (LDServerSDK_Start(client, maxwait, &initialized_successfully)) {
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
