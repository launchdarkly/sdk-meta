---
id: android-client-sdk/sdk-docs/evaluate-a-flag-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
description: "Kotlin in section \"Evaluate a flag\""
validation:
  scaffold: android-client-sdk/scaffolds/android-syntax-only
---

```kotlin
val showFeature: Boolean = client.boolVariation(flagKey, true)
if (showFeature) {
    // Application code to show the feature
}
else {
    // The code to run if the feature is off
}
```
