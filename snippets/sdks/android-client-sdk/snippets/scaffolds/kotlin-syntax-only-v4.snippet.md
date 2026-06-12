---
id: android-client-sdk/scaffolds/kotlin-syntax-only-v4
sdk: android-client-sdk
kind: scaffold
lang: kotlin
file: app/src/main/java/com/launchdarkly/hello_android/Snippet.kt
description: |
  Parse-and-type-check validator for Android Kotlin doc fragments
  that target the v4.x API surface (`LDConfig.Builder()` with no
  arguments — v5 made the `AutoEnvAttributes` constructor argument
  mandatory, so v4-era fragments cannot compile against the v5 aar
  the `android-client` validator container pre-bakes).

  Unlike the Java sibling (`java-syntax-only-v4`, which routes
  through the `jvm` validator because a Java file can only carry one
  top-level class), this scaffold stays in the `android-client`
  container in `SNIPPET_CHECK=parse` mode: Kotlin permits multiple
  top-level classes per file, and same-file declarations take
  resolution priority over star imports, so file-scope stub classes
  named `LDConfig` / `LDClient` shadow nothing and need no aar at
  all. Only the v4 surface the doc fragments call is declared; the
  shared `com.launchdarkly.sdk.LDContext` type is real (it ships
  inside the v5 aar's java-sdk-common dependency and is
  version-stable across v4/v5).

  File-scope `application` / `context` stubs mirror the ambient
  names the doc fragments assume an Activity host provides. The
  wrappee body is spliced inside an unreachable `if (false)` block
  in a never-invoked function.
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
import com.launchdarkly.sdk.LDContext

// Stub of the v4 android LDConfig builder chain. Only the members
// the doc fragments call are declared; each returns a stub so the
// chained-builder shape type-checks.
class LDConfig {
    class Builder {
        fun mobileKey(key: String): Builder = this
        fun offline(offline: Boolean): Builder = this
        fun build(): LDConfig = LDConfig()
    }
}

// Stub of the v4 android LDClient surface the doc fragments touch.
class LDClient private constructor() {
    fun setOffline() {}
    companion object {
        fun init(
            application: Application,
            config: LDConfig,
            context: LDContext,
            startWaitSeconds: Int
        ): LDClient = LDClient()
    }
}

// Ambient names the doc fragments assume an Activity host provides.
// Never read at runtime (the body below is unreachable).
@Suppress("UNUSED")
val application: Application get() = TODO()
@Suppress("UNUSED")
val context: LDContext get() = TODO()

@Suppress("UNUSED_VARIABLE", "UNREACHABLE_CODE", "unused")
private fun _wrappee() {
    if (false) {
{{ body }}
    }
}
```
