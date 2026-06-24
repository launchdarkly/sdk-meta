---
id: cpp-server-sdk/sdk-docs/features/filedata/flags-from-files-c-sdk-v2
sdk: cpp-server-sdk
kind: reference
lang: c
description: File data source configuration example for the v2.x C server SDK.
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only-v2-c

---

```c
#include <launchdarkly/integrations/file_data.h>

const char *filenames[2] = {
    "file1.json",
    "file2.json"
};
LDConfigSetDataSource(config, LDFileDataInit(2, filenames));
LDConfigSetSendEvents(config, LDBooleanFalse);

// Call LDClientInit with config as usual.
```
