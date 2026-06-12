---
id: android-client-sdk/scaffolds/kotlin-syntax-only
sdk: android-client-sdk
kind: scaffold
lang: kotlin
file: app/src/main/java/com/launchdarkly/hello_android/Snippet.kt
description: |
  Parse-and-type-check validator for Android Kotlin doc fragments.

  Routes through the `android-client` validator container in
  `SNIPPET_CHECK=parse` mode: the harness stages this file into the
  pre-baked hello-android gradle project alongside the existing
  MainApplication / MainActivity sources, then runs
  `./gradlew compileDebugKotlin`. A passing compile means the body
  parses AND type-checks against the real
  `com.launchdarkly:launchdarkly-android-client-sdk` aar from
  Google's Maven plus the AndroidX runtime — neither of which the
  `jvm` validator's Java + Maven path can resolve.

  The wrappee body is spliced inside `BaseApplication.onCreate()`'s
  unreachable `if (false)` block so unresolved caller surfaces
  (an Activity host's lifecycle, etc.) and Kotlin's
  `this@BaseApplication` labeled-this both have somewhere legal to
  land. Bodies declaring `class MainActivity : AppCompatActivity()`
  end up as local classes inside onCreate — Kotlin permits that
  shape.

  Bodies with `import …` lines from inside the body block are
  handled by the harness's existing Python import-lift pre-step:
  Kotlin only allows imports between `package` and the first
  top-level declaration, so they're hoisted to file scope before
  kotlinc runs.

  File-scope stubs cover caller-supplied helpers the experimentation
  bodies reference (`applyVariant(_:)`); they're declared at module
  level here so the body's call sites resolve at compile time.
  Never invoked.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: android-client
  entrypoint: app/src/main/java/com/launchdarkly/hello_android/Snippet.kt
  env:
    SNIPPET_CHECK: parse
---

```kotlin
package com.launchdarkly.hello_android

import android.app.Application
import android.os.Bundle
import androidx.appcompat.app.AppCompatActivity
// Wildcard imports cover the doc fragments' references to `LDConfig`,
// `LDContext`, `Components`, `Plugin`, `Observability`, etc. without
// per-snippet placeholder boilerplate.
import com.launchdarkly.sdk.*
import com.launchdarkly.sdk.android.*
import com.launchdarkly.sdk.android.integrations.*
import com.launchdarkly.observability.plugin.*
// `AutoEnvAttributes` is a nested enum at
// `LDConfig.Builder.AutoEnvAttributes` (Java class name
// `LDConfig$Builder$AutoEnvAttributes`). Package wildcard imports
// don't reach nested types, so it needs an explicit import for the
// v5-x init body's `AutoEnvAttributes.Enabled` reference.
import com.launchdarkly.sdk.android.LDConfig.Builder.AutoEnvAttributes

// File-scope stubs so wrappee bodies that reference caller-supplied
// helpers (e.g. `applyVariant(variant)`) type-check. Never invoked.
fun applyVariant(variant: String) {}

// File-scope stubs for the bare `client` and `flagKey` references
// the evaluate-a-flag-kotlin body assumes are already in scope. The
// scaffold's `if (false)` block guarantees these are never read at
// runtime, so initializing `client` via TODO() (which returns
// Nothing) is safe — top-level vals are evaluated lazily-on-read in
// most cases but to be safe the body is unreachable anyway.
@Suppress("UNUSED")
val client: LDClient get() = TODO()
@Suppress("UNUSED")
val flagKey: String = ""
// Init fragments pass the ambient `application` instance the docs
// assume an Activity host provides.
@Suppress("UNUSED")
val application: Application get() = TODO()

@Suppress("UNUSED_VARIABLE", "UNREACHABLE_CODE")
class BaseApplication : Application() {
    override fun onCreate() {
        super.onCreate()
        if (false) {
{{ body }}
        }
    }
}
```
