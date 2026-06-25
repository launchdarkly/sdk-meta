---
id: android-client-sdk/sdk-docs/features/localstorage/localstorage-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
description: Local storage caching example for Android (Kotlin).
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-syntax-only
---

```kotlin
val ldConfig = LDConfig.Builder(AutoEnvAttributes.Enabled)
    .mobileKey("example-mobile-key")
    // Local storage is enabled by default
    // You can optionally configure the maximum number of cached contexts (default is 5)
    .maxCachedContexts(3)
    .build()

val context = LDContext.create("example-context-key")

val client = LDClient.init(application, ldConfig, context, 0)
```
