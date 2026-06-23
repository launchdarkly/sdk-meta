---
id: android-client-sdk/sdk-docs/features/privateattrs/config-v5-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
description: Private attribute configuration for Android SDK v5.x (Kotlin).
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-syntax-only

---

```kotlin
// All attributes marked private
var ldConfig = LDConfig.Builder(AutoEnvAttributes.Enabled)
    .mobileKey("example-mobile-key")
    .events(
        Components.sendEvents()
          .allAttributesPrivate(true)
     )
     .build()

// Two attributes marked private
ldConfig = LDConfig.Builder(AutoEnvAttributes.Enabled)
    .mobileKey("example-mobile-key")
    .events(
        Components.sendEvents()
            .privateAttributes("name", "group")
    )
    .build();
```
