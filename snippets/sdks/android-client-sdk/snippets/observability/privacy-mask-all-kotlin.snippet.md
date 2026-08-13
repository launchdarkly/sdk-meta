---
id: android-client-sdk/observability/privacy-mask-all-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
file: app/src/main/java/com/launchdarkly/hello_android/Snippet.kt
description: "Android observability plugin: Privacy profile masking all (Kotlin)"
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-observability
---

```kotlin
privacyProfile = PrivacyProfile(
  maskTextInputs = true,
  maskText = true,
  maskBySemanticsKeywords = true
)
```
