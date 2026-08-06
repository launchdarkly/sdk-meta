---
id: android-client-sdk/observability/flush-session-replay-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
file: app/src/main/java/com/launchdarkly/hello_android/Snippet.kt
description: "Android observability plugin: Flush session replay (Kotlin)"
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-observability
---

```kotlin
import com.launchdarkly.observability.sdk.LDReplay

// Immediately export any queued replay events
LDReplay.flush()
```
