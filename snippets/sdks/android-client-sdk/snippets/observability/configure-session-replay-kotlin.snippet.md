---
id: android-client-sdk/observability/configure-session-replay-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
file: app/src/main/java/com/launchdarkly/hello_android/Snippet.kt
description: "Android observability plugin: Configure session replay (Kotlin)"
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-observability
---

```kotlin
import com.launchdarkly.observability.plugin.Observability
import com.launchdarkly.observability.replay.plugin.SessionReplay

val mobileKey = "example-mobile-key"

val ldConfig = LDConfig.Builder(AutoEnvAttributes.Enabled)
    .mobileKey(mobileKey)
    .plugins(
      Components.plugins().setPlugins(
        listOf(
          Observability(this@BaseApplication, mobileKey),
          SessionReplay()  // depends on Observability being present first
        )
      )
    )
    .build()
```
