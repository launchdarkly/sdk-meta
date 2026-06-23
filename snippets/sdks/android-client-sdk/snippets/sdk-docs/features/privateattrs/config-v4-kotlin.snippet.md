---
id: android-client-sdk/sdk-docs/features/privateattrs/config-v4-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
description: Private attribute configuration for Android SDK v4.x (Kotlin).
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-syntax-only-v4

---

```kotlin
// All attributes marked private
var config = LDConfig.Builder()
    .mobileKey("example-mobile-key")
    .events(
        Components.sendEvents()
            .allAttributesPrivate(true)
    )
    .build();

// Two attributes marked private
config = LDConfig.Builder()
    .mobileKey("example-mobile-key")
    .events(
        Components.sendEvents()
            .privateAttributes("name", "group")
    )
    .build();
```
