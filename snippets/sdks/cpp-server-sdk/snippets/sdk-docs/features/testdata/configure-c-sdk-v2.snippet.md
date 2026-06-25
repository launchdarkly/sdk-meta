---
id: cpp-server-sdk/sdk-docs/features/testdata/configure-c-sdk-v2
sdk: cpp-server-sdk
kind: reference
lang: c
description: Test data source configuration for the C server SDK v2.x (native).
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only-v2-c
---

```c
#include <launchdarkly/integrations/test_data.h>

struct LDTestData *td = LDTestDataInit();

LDConfigSetDataSource(config, LDTestDataCreateDataSource(td));

// Call LDClientInit with config as usual.
```
