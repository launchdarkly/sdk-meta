---
id: android-client-sdk/observability/initialize-plugins-after-client-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
file: app/src/main/java/com/launchdarkly/hello_android/Snippet.kt
description: "Android observability plugin: Initialize the plugins after the SDK client (Kotlin)"
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-observability
---

```kotlin
import com.launchdarkly.observability.plugin.Observability
import com.launchdarkly.observability.replay.plugin.SessionReplay
import com.launchdarkly.observability.replay.ReplayOptions

val mobileKey = "example-mobile-key"

val ldConfig = LDConfig.Builder(AutoEnvAttributes.Enabled)
    .mobileKey(mobileKey)
    .plugins(
      Components.plugins().setPlugins(
        listOf(
          Observability(this@BaseApplication, mobileKey),
          SessionReplay(
            options = ReplayOptions(
              enabled = false // Don't start recording automatically
            )
          )
        )
      )
    )
    .build()

val context = LDContext.create("example-context-key")
val client = LDClient.init(this@BaseApplication, ldConfig, context, 0)
```
