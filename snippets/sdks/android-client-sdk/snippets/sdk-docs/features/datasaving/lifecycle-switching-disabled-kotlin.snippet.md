---
id: android-client-sdk/sdk-docs/features/datasaving/lifecycle-switching-disabled-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
description: Disable lifecycle-driven mode switching while keeping network-driven switching for Android (Kotlin).
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-syntax-only
---

```kotlin
val config = LDConfig.Builder(AutoEnvAttributes.Enabled)
    .mobileKey("example-mobile-key")
    .dataSystem(
        Components.dataSystem()
            .automaticModeSwitching(
                DataSystemComponents.automaticModeSwitching()
                    .lifecycle(false)
                    .network(true)
                    .build()))
    .build()
```
