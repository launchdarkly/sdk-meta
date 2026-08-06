---
id: android-client-sdk/observability/distributed-tracing-propagate-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
file: app/src/main/java/com/launchdarkly/hello_android/Snippet.kt
description: "Android observability plugin: Distributed tracing, propagate headers (Kotlin)"
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-observability
---

```kotlin
import io.opentelemetry.api.trace.propagation.W3CTraceContextPropagator
import io.opentelemetry.context.Context
import okhttp3.Request

val propagator = W3CTraceContextPropagator.getInstance()
val requestBuilder = Request.Builder().url(url)

propagator.inject(
    Context.current(),
    requestBuilder
) { carrier, key, value -> carrier?.addHeader(key, value) }

val request = requestBuilder.build()
```
