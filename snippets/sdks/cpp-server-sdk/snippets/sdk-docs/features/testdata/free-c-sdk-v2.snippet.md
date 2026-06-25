---
id: cpp-server-sdk/sdk-docs/features/testdata/free-c-sdk-v2
sdk: cpp-server-sdk
kind: reference
lang: c
description: Freeing the test data source for the C server SDK v2.x (native).
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only-v2-c
---

```c
// After LDClientClose:
LDTestDataFree(td);
```
