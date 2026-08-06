---
id: android-client-sdk/observability/start-session-replay-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
file: app/src/main/java/com/launchdarkly/hello_android/Snippet.kt
description: "Android observability plugin: Start session replay (Kotlin)"
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-observability
---

```kotlin
import com.launchdarkly.observability.sdk.LDReplay

// Start recording after user consent or feature flag check
LDReplay.start()
```
