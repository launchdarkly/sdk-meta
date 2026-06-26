---
id: android-client-sdk/sdk-docs/features/datasaving/disable-mode-switching-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
description: Disable automatic mode switching entirely for Android (Kotlin).
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-syntax-only
---

```kotlin
val config = LDConfig.Builder(AutoEnvAttributes.Enabled)
    .mobileKey("example-mobile-key")
    .dataSystem(
        Components.dataSystem()
            .automaticModeSwitching(AutomaticModeSwitchingConfig.disabled())
            .foregroundConnectionMode(ConnectionMode.STREAMING))
    .build()
```
