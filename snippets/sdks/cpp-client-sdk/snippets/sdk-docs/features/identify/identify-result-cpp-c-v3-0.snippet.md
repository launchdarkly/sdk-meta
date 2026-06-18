---
id: cpp-client-sdk/sdk-docs/features/identify/identify-result-cpp-c-v3-0
sdk: cpp-client-sdk
kind: reference
lang: c
description: Examining the identify result for the C++ client SDK v3.0 (C binding).
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only
---

```c
bool identified_successfully;
if (LDClientSDK_Identify(client, updated_context, maxwait, &identified_successfully)) {
  /* The client's attempt to identify succeeded or failed in the specified amount of time. */
  if (identified_successfully) {
    /* Identification succeeded. */
  } else {
    /* Identification failed. */
  }
} else {
   /* The specified timeout was reached, but the client is still identifying. */
}
```
