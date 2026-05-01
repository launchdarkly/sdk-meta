---
id: android-client-sdk/sdk-docs/migration-2-to-3-track-methods-3-0-syntax-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
description: "3.0 syntax (Kotlin) in section \"Track methods\""
---

```kotlin
client.trackData("dataEvent", LDValue.of(4))
client.trackMetric("metricEvent", LDValue.ofNull(), 5.5)
client.trackMetric("bothEvent", LDValue.of("tag"), 3.5)
```
