---
id: android-client-sdk/sdk-docs/evaluate-a-flag-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
description: "Kotlin in section \"Evaluate a flag\""
# Bucket C: jvm validator pulls launchdarkly-java-server-sdk, not the android-client SDK (which lives in Google Maven as an aar). See _sdk-docs-port-notes.md.
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-syntax-only
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
