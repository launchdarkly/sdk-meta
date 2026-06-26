---
id: java-server-sdk/sdk-docs/features/datasaving/file-bootstrap
sdk: java-server-sdk
kind: reference
lang: java
description: Data saving mode with file-based bootstrap and live updates for Java.
validation:
  scaffold: java-server-sdk/scaffolds/java-syntax-only
---

```java
import com.launchdarkly.sdk.server.*;
import com.launchdarkly.sdk.server.integrations.*;

LDConfig config = new LDConfig.Builder()
    .dataSystem(
        Components.dataSystem().custom()
            .initializers(
                FileData.initializer().filePaths("flags.json"),
                DataSystemComponents.pollingInitializer()
            )
            .synchronizers(
                DataSystemComponents.streamingSynchronizer(),
                DataSystemComponents.pollingSynchronizer()
            )
    )
    .build();

LDClient client = new LDClient("YOUR_SDK_KEY", config);
```
