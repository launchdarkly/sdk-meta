---
id: java-server-sdk/sdk-docs/features/relay-proxy-config/daemon-mode/daemon-mode-datasystem
sdk: java-server-sdk
kind: reference
lang: java
description: Daemon mode DataSystem configuration example for Java.
validation:
  scaffold: java-server-sdk/scaffolds/java-syntax-only
---

```java
import com.launchdarkly.sdk.server.*;

LDConfig config = new LDConfig.Builder()
    .dataSystem(
        Components.dataSystem().daemon(
            Components.persistentDataStore(Redis.dataStore())
        )
    )
    .build();

LDClient client = new LDClient("YOUR_SDK_KEY", config);
```
