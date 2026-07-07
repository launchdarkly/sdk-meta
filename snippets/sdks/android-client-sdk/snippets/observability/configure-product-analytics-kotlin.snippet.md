---
id: android-client-sdk/observability/configure-product-analytics-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
file: app/src/main/java/com/launchdarkly/hello_android/Snippet.kt
description: "Android observability plugin: Configure product analytics event collection (Kotlin)"
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-observability
---

```kotlin
ObservabilityOptions(
  analytics = ObservabilityOptions.Analytics(
    taps = true,
    trackEvents = true,
    screenViews = true,
    appLifecycle = true,
    appLaunch = true,
  )
)
```
