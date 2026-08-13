---
id: android-client-sdk/observability/privacy-mask-semantics-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
file: app/src/main/java/com/launchdarkly/hello_android/Snippet.kt
description: "Android observability plugin: Privacy profile masking by semantics (Kotlin)"
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-observability
---

```kotlin
privacyProfile = PrivacyProfile(
  maskTextInputs = true,
  maskText = false,
  maskBySemanticsKeywords = true
)
```
