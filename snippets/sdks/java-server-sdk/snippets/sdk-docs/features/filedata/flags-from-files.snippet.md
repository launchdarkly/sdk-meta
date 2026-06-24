---
id: java-server-sdk/sdk-docs/features/filedata/flags-from-files
sdk: java-server-sdk
kind: reference
lang: java
description: File data source configuration example for Java.
validation:
  scaffold: java-server-sdk/scaffolds/java-syntax-only

---

```java
import com.launchdarkly.sdk.server.*;
import com.launchdarkly.sdk.server.integrations.*;

LDConfig config = new LDConfig.Builder()
  .dataSource(
    FileData.dataSource()
      .filePaths("file1.json", "file2.json")
      .autoUpdate(true)
  )
  .events(Components.noEvents())
  .build();

LDClient client = new LDClient("sdk key", config);
```
