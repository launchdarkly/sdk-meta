---
id: android-client-sdk/sdk-docs/initialize-the-client-android-sdk-v5-x-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
description: "Android SDK v5.x (Kotlin) in section \"Initialize the client\""
# Bucket C: jvm validator pulls launchdarkly-java-server-sdk, not the android-client SDK (which lives in Google Maven as an aar). See _sdk-docs-port-notes.md.
---

```kotlin
val ldConfig = LDConfig.Builder(AutoEnvAttributes.Enabled)
    .mobileKey("example-mobile-key")
    // optional observability plugin, requires LaunchDarkly Android Client SDK v5.9+
    .plugins(Components.plugins().setPlugins(
      listOf(Observability(this@BaseApplication, "example-mobile-key"))
    ))
    // other options
    .build()

// You'll need this context later, but you can ignore it for now.
val context = LDContext.create("example-context-key")

val client: LDClient = LDClient.init(this@BaseApplication, ldConfig, context, 0)
```
