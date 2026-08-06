---
id: android-client-sdk/observability/advanced-configuration-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
file: app/src/main/java/com/launchdarkly/hello_android/Snippet.kt
description: "Android observability plugin: Advanced configuration options (Kotlin)"
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-observability
---

```kotlin
val mobileKey = "example-mobile-key"

val ldConfig = LDConfig.Builder(AutoEnvAttributes.Enabled)
    .mobileKey(mobileKey)
    .plugins(
      Components.plugins().setPlugins(
        listOf(
          Observability(
            this@BaseApplication,
            mobileKey,
            ObservabilityOptions(
              serviceName = "my-android-app",
              serviceVersion = "1.0.0",
              debug = true,
              logsApiLevel = ObservabilityOptions.LogLevel.WARN,
              tracesApi = ObservabilityOptions.TracesApi(includeErrors = true, includeSpans = false),
              metricsApi = ObservabilityOptions.MetricsApi.disabled(),
              instrumentations = ObservabilityOptions.Instrumentations(
                crashReporting = false,
                launchTime = true,
                userTaps = true,
                screens = true
              ),
              resourceAttributes = Attributes.of(
                AttributeKey.stringKey("environment"), "production",
                AttributeKey.stringKey("team"), "mobile"
              ),
              customHeaders = mapOf(
                "X-Custom-Header" to "custom-value"
              )
            )
          )
        )
      )
    )
    .build()
```
