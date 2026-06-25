---
id: android-client-sdk/sdk-docs/features/flagchanges/all-flags-listener-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
description: All-flags update listener example for Android (Kotlin).
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-syntax-only

---

```kotlin
val listener = LDAllFlagsListener {
    // Get new values for flag keys (from implicit "it" list variable) or other operations
}

// register all flags listener
LDClient.get().registerAllFlagsListener(listener)
// when done with all flags listener it should be unregistered
LDClient.get().unregisterAllFlagsListener(listener)
```
