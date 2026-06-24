---
id: android-client-sdk/sdk-docs/features/flagchanges/flag-changes-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
description: Flag change listener registration for Android (Kotlin).
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-syntax-only

---

```kotlin
val flagKey = "yourFlagKey"

val listener = FeatureFlagChangeListener {
    val newValue = LDClient.get().boolVariation(flagKey, false)
}

LDClient.get().registerFeatureFlagListener(flagKey, listener)
```
