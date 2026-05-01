---
id: android-client-sdk/sdk-docs/initialize-the-client-android-sdk-v4-x-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
description: "Android SDK v4.x (Kotlin) in section \"Initialize the client\""
---

```kotlin
val ldConfig = LDConfig.Builder()
    .mobileKey("example-mobile-key")
    .build()

// You'll need this context later, but you can ignore it for now.
val context = LDContext.create("example-context-key")

val client: LDClient = LDClient.init(this@BaseApplication, ldConfig, context, 0)
```
