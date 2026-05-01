---
id: android-client-sdk/sdk-docs/migration-2-to-3-track-methods-2-x-syntax-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
description: "2.x syntax (Kotlin) in section \"Track methods\""
---

```kotlin
client.track("dataEvent", JsonPrimitive(4))
client.track("metricEvent", null, 5.5)
client.track("bothEvent", JsonPrimitive("tag"), 3.5)
```
