---
id: android-client-sdk/scaffolds/kotlin-observability
sdk: android-client-sdk
kind: scaffold
lang: kotlin
file: app/src/main/java/com/launchdarkly/hello_android/Snippet.kt
description: |
  Type-check validator for Android observability-plugin Kotlin doc
  fragments. Same `android-client` container / `compileDebugKotlin`
  path as `kotlin-syntax-only`, but with the observability plugin's
  packages and its OpenTelemetry dependency imported at file scope, so
  the body type-checks against the real
  `com.launchdarkly:launchdarkly-observability-android` aar (plugin,
  api, sdk, and session-replay packages) plus the OTel API.

  The body is spliced inside `BaseApplication.onCreate()`'s unreachable
  `if (false)` block, so `this@BaseApplication` and the SDK/plugin
  init surfaces have a legal home; the harness's import-lift pre-step
  hoists any `import …` lines the fragment carries up to file scope.
  File-scope stubs stand in for the ambient values the fragments
  assume from surrounding application code (`mobileKey`, a `text`
  log message, a request `url`, a `privacyProfile` the option-only
  fragments assign).
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, type-checked against the plugin.
validation:
  runtime: android-client
  entrypoint: app/src/main/java/com/launchdarkly/hello_android/Snippet.kt
  env:
    SNIPPET_CHECK: parse
---

```kotlin
package com.launchdarkly.hello_android

import android.app.Application
import android.app.Activity
import android.os.Bundle
import androidx.appcompat.app.AppCompatActivity
import com.launchdarkly.sdk.*
import com.launchdarkly.sdk.android.*
import com.launchdarkly.sdk.android.integrations.*
import com.launchdarkly.sdk.android.LDConfig.Builder.AutoEnvAttributes
import com.launchdarkly.observability.plugin.*
import com.launchdarkly.observability.api.*
import com.launchdarkly.observability.sdk.*
import com.launchdarkly.observability.replay.*
import com.launchdarkly.observability.replay.plugin.*
// Explicit OTel imports (not wildcards): the observability AAR ships a
// `com.launchdarkly.observability.replay.Attributes`, so a
// `io.opentelemetry.api.common.*` wildcard would let the replay type
// shadow OTel's `Attributes` for the unqualified name the doc fragments
// use. Explicit single imports outrank star imports and disambiguate.
import io.opentelemetry.api.common.Attributes
import io.opentelemetry.api.common.AttributeKey
import io.opentelemetry.api.trace.Span
import io.opentelemetry.api.trace.SpanContext
import io.opentelemetry.api.trace.propagation.W3CTraceContextPropagator
import io.opentelemetry.api.logs.Severity
import io.opentelemetry.context.Context
import io.opentelemetry.context.Scope
import kotlinx.coroutines.*
import okhttp3.Request

// File-scope stubs for the ambient values observability fragments assume
// from surrounding application code. Never read at runtime — the body is
// spliced into an unreachable `if (false)` block.
@Suppress("UNUSED")
val mobileKey: String = ""
@Suppress("UNUSED")
val text: String = ""
@Suppress("UNUSED")
val url: String = ""
// The privacy-options fragments show only a `privacyProfile = PrivacyProfile(...)`
// assignment; provide the target so they type-check against the real type.
@Suppress("UNUSED")
var privacyProfile: PrivacyProfile? = null
// Stand-in for the application work the tracing fragments wrap in a span.
@Suppress("UNUSED")
fun performDatabaseQuery() {}

@Suppress("UNUSED_VARIABLE", "UNREACHABLE_CODE", "UNUSED_EXPRESSION")
class BaseApplication : Application() {
    override fun onCreate() {
        super.onCreate()
        if (false) {
{{ body }}
        }
    }
}
```
