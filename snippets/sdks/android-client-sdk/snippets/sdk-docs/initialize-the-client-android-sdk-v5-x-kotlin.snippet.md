---
id: android-client-sdk/sdk-docs/initialize-the-client-android-sdk-v5-x-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
description: "Android SDK v5.x (Kotlin) in section \"Initialize the client\""
---

```kotlin
val ldConfig = LDConfig.Builder(AutoEnvAttributes.Enabled)
    .mobileKey("example-mobile-key")
    // optional observability plugin, requires LaunchDarkly Android Client SDK v5.9+
    .plugins(Components.plugins().setPlugins(
      listOf(Observability(this@BaseApplication))
    ))
    // other options
    .build()

// You'll need this context later, but you can ignore it for now.
val context = LDContext.create("example-context-key")

val client: LDClient = LDClient.init(this@BaseApplication, ldConfig, context, 0)
```
