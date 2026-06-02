---
id: android-client-sdk/sdk-docs/initialize-the-client-android-sdk-v4-x-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
description: "Android SDK v4.x (Kotlin) in section \"Initialize the client\""
# TODO(validator): same version-pinned issue as the v4-x-java
# sibling — `LDConfig.Builder()` is 0-arg in v4 but requires
# `AutoEnvAttributes` in v5+. The kotlin-syntax-only scaffold
# compiles against the v5.x AAR. Needs the same android-client-v4
# version-pinned validator.
---

```kotlin
val ldConfig = LDConfig.Builder()
    .mobileKey("example-mobile-key")
    .build()

// You'll need this context later, but you can ignore it for now.
val context = LDContext.create("example-context-key")

val client: LDClient = LDClient.init(this@BaseApplication, ldConfig, context, 0)
```
