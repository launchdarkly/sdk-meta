---
id: android-client-sdk/observability/privacy-mask-none-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
file: app/src/main/java/com/launchdarkly/hello_android/Snippet.kt
description: "Android observability plugin: Privacy profile masking none (Kotlin)"
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-observability
---

```kotlin
privacyProfile = PrivacyProfile(
  maskTextInputs = false,
  maskText = false,
  maskBySemanticsKeywords = false
)
```
