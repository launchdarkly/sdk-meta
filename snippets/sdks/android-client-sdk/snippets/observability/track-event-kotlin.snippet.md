---
id: android-client-sdk/observability/track-event-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
file: app/src/main/java/com/launchdarkly/hello_android/Snippet.kt
description: "Android observability plugin: Track a custom product analytics event (Kotlin)"
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-observability
---

```kotlin
// Track an event with properties and an optional metric value
LDObserve.track(
    key = "purchase_completed",
    properties = mapOf(
        "product_id" to "SKU-123",
        "price" to 29.99
    ),
    metricValue = 29.99
)

// Track an event with no properties
LDObserve.track(key = "button_tapped")
```
