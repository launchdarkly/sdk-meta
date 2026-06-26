---
id: java-server-sdk/sdk-docs/features/bigsegments/big-segments
sdk: java-server-sdk
kind: reference
lang: java
description: Big segments Redis store configuration example for Java.
validation:
  scaffold: java-server-sdk/scaffolds/java-syntax-only

---

```java
import com.launchdarkly.sdk.server.*;
import com.launchdarkly.sdk.server.integrations.*;

import java.net.URI;
import java.time.Duration;

LDConfig config = new LDConfig.Builder()
    .bigSegments(
        Components.bigSegments(
            Redis.bigSegmentStore()
                .uri(URI.create("redis://your-redis:6379"))
                .prefix("example-client-side-id")
        )
        .userCacheSize(2000)
        .userCacheTime(Duration.ofSeconds(30))
    )
    .build();
LDClient client = new LDClient("YOUR_SDK_KEY", config);
```
