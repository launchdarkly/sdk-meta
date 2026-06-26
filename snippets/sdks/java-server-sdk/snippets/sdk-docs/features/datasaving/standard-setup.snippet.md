---
id: java-server-sdk/sdk-docs/features/datasaving/standard-setup
sdk: java-server-sdk
kind: reference
lang: java
description: Data saving mode standard setup for Java.
validation:
  scaffold: java-server-sdk/scaffolds/java-syntax-only
---

```java
import com.launchdarkly.sdk.server.*;

LDConfig config = new LDConfig.Builder()
    .dataSystem(Components.dataSystem().defaultMode())
    .build();

LDClient client = new LDClient("YOUR_SDK_KEY", config);
```
