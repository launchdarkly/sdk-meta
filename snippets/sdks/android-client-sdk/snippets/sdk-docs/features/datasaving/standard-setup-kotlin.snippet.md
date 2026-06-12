---
id: android-client-sdk/sdk-docs/features/datasaving/standard-setup-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
description: Data saving mode standard setup for Android (Kotlin).
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-syntax-only
---

```kotlin
val config = LDConfig.Builder(AutoEnvAttributes.Enabled)
    .mobileKey("example-mobile-key")
    .dataSystem(Components.dataSystem())
    .build()
```
