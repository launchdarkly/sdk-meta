---
id: cpp-client-sdk/sdk-docs/features/flush/flush-result-cpp-c-v3-0
sdk: cpp-client-sdk
kind: reference
lang: cpp
description: Flush result example for C++ (client-side) SDK v3.0 (C binding).
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only

---

```cpp
/* LDClientSDK_Flush returns immediately; event delivery proceeds in
   the background. The second parameter must be LD_NONBLOCKING, and the
   call does not report a result. */
LDClientSDK_Flush(client, LD_NONBLOCKING);
```
