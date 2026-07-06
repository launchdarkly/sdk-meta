---
id: android-client-sdk/observability/initialize-the-client-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
file: app/src/main/java/com/launchdarkly/hello_android/Snippet.kt
description: "Android observability plugin in section \"Initialize the client\" (Kotlin)"
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-observability
---

```kotlin
val mobileKey = "example-mobile-key"

val ldConfig = LDConfig.Builder(AutoEnvAttributes.Enabled)
    .mobileKey(mobileKey)
    .plugins(Components.plugins().setPlugins(
      listOf(Observability(this@BaseApplication, mobileKey))
    ))
    // other options
    .build()

// You'll need this context later, but you can ignore it for now.
val context = LDContext.create("example-context-key")

val client: LDClient = LDClient.init(this@BaseApplication, ldConfig, context, 0)
```
