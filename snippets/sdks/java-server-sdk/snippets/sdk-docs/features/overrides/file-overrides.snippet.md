---
id: java-server-sdk/sdk-docs/features/overrides/file-overrides
sdk: java-server-sdk
kind: reference
lang: java
description: File-based flag override source configuration example for Java.
validation:
  scaffold: java-server-sdk/scaffolds/java-syntax-only
---

```java
import com.launchdarkly.sdk.server.*;
import com.launchdarkly.sdk.server.integrations.*;

LDConfig config = new LDConfig.Builder()
    .dataSystem(
        Components.dataSystem().defaultMode()
            .overrides(FileOverrides.source().filePaths("/etc/launchdarkly/overrides.json"))
    )
    .build();

LDClient client = new LDClient("YOUR_SDK_KEY", config);
```
