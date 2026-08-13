---
id: android-client-sdk/observability/track-screen-view-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
file: app/src/main/java/com/launchdarkly/hello_android/Snippet.kt
description: "Android observability plugin: Track a screen view (Kotlin)"
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-observability
---

```kotlin
// Convenience — name only
LDObserve.trackScreenView(name = "ProductDetail")

// With category
LDObserve.trackScreenView(name = "Checkout", category = "purchase")

// Full details with custom properties
LDObserve.trackScreenView(
    name = "ProductDetail",
    screenClass = "ProductDetailActivity",
    screenId = "product-123",
    category = "browsing",
    properties = mapOf(
        "product_id" to "SKU-123",
        "source" to "search_results"
    )
)
```
