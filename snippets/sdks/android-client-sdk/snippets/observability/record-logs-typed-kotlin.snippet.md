---
id: android-client-sdk/observability/record-logs-typed-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
file: app/src/main/java/com/launchdarkly/hello_android/Snippet.kt
description: "Android observability plugin: Record custom logs with OTel-typed attributes (Kotlin)"
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-observability
---

```kotlin
// Use OTel-typed attributes when exact OpenTelemetry typing is required
LDObserve.recordLog(
    message = "Authentication completed",
    severity = Severity.INFO,
    attributes = Attributes.of(
        AttributeKey.stringKey("user_id"), "12345",
        AttributeKey.stringKey("action"), "login"
    )
)
```
