---
id: android-client-sdk/sdk-docs/migration-2-to-3-jackson-kotlin
sdk: android-client-sdk
kind: reference
lang: java
description: "Kotlin in section \"Jackson\""
---

```java
import com.google.jackson.databind.*
import com.launchdarkly.sdk.json.LDJackson

val mapper: ObjectMapper = ObjectMapper()
mapper.registerModule(LDJackson.module())
```
