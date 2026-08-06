---
id: android-client-sdk/observability/record-logs-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
file: app/src/main/java/com/launchdarkly/hello_android/Snippet.kt
description: "Android observability plugin: Record custom logs (Kotlin)"
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-observability
---

```kotlin
// Record a basic log message
LDObserve.recordLog(
    message = "User login successful",
    severity = Severity.INFO
)

// Record logs with custom properties
LDObserve.recordLog(
    message = "Authentication completed",
    severity = Severity.INFO,
    properties = mapOf(
        "user_id" to "12345",
        "action" to "login"
    )
)
```
