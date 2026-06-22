---
id: java-server-sdk/sdk-docs/features/logging/logging
sdk: java-server-sdk
kind: reference
lang: java
description: Logging destination and level configuration example for Java SDK v5.10.x and later.
validation:
  scaffold: java-server-sdk/scaffolds/java-syntax-only

---

```java
import com.launchdarkly.logging.*;
import com.launchdarkly.sdk.server.*;

LDConfig config = new LDConfig.Builder()
  .logging(
    Components.logging(Logs.toStream(System.out)).level(LDLogLevel.DEBUG)
  )
  .build();
```
