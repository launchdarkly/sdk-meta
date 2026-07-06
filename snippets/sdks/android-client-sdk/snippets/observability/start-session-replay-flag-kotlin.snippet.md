---
id: android-client-sdk/observability/start-session-replay-flag-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
file: app/src/main/java/com/launchdarkly/hello_android/Snippet.kt
description: "Android observability plugin: Start session replay from a flag (Kotlin)"
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-observability
---

```kotlin
import com.launchdarkly.observability.sdk.LDReplay

// Start recording based on a feature flag
val replayEnabled = LDClient.get().boolVariation("enable-session-replay", false)
if (replayEnabled) {
    LDReplay.start()
}
```
