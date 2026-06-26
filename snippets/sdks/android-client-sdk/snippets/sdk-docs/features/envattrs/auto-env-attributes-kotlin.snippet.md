---
id: android-client-sdk/sdk-docs/features/envattrs/auto-env-attributes-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
description: Automatic environment attributes configuration for Android (Kotlin).
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-syntax-only

---

```kotlin
val ldConfig = LDConfig.Builder(AutoEnvAttributes.Enabled)
    .mobileKey("example-mobile-key")
    .build()
```
