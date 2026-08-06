---
id: android-client-sdk/observability/session-replay-privacy-options-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
file: app/src/main/java/com/launchdarkly/hello_android/Snippet.kt
description: "Android observability plugin: Privacy options (Kotlin)"
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-observability
---

```kotlin
import com.launchdarkly.observability.plugin.Observability
import com.launchdarkly.observability.replay.plugin.SessionReplay
import com.launchdarkly.observability.replay.ReplayOptions
import com.launchdarkly.observability.replay.PrivacyProfile

val mobileKey = "example-mobile-key"

val ldConfig = LDConfig.Builder(AutoEnvAttributes.Enabled)
    .mobileKey(mobileKey)
    .plugins(
      Components.plugins().setPlugins(
        listOf(
          Observability(this@BaseApplication, mobileKey),
          SessionReplay(
            options = ReplayOptions(
              privacyProfile = PrivacyProfile(
                maskTextInputs = true,
                maskText = false
              )
            )
          )
        )
      )
    )
    .build()
```
