---
id: cpp-server-sdk/sdk-docs/features/testdata/set-flag-value-c-sdk-v2
sdk: cpp-server-sdk
kind: reference
lang: c
description: Setting a test data flag to a specific value for the C server SDK v2.x (native).
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only-v2-c
---

```c
struct LDFlagBuilder *flag = LDTestDataFlag(td, "example-flag-key");
LDFlagBuilderVariationForAllUsersBoolean(flag, LDBooleanTrue);
LDTestDataUpdate(td, flag);
```
