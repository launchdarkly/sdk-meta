---
id: java-server-sdk/sdk-docs/migration-4-to-5-jackson-java
sdk: java-server-sdk
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
