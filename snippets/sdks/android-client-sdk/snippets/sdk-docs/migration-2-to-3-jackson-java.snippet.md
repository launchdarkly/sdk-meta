---
id: android-client-sdk/sdk-docs/migration-2-to-3-jackson-java
sdk: android-client-sdk
kind: reference
lang: java
description: "Java in section \"Jackson\""
---

```java
import com.google.jackson.databind.*;
import com.launchdarkly.sdk.json.LDJackson;

ObjectMapper mapper = new ObjectMapper();
mapper.registerModule(LDJackson.module());
```
