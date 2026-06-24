---
id: android-client-sdk/scaffolds/kotlin-syntax-only-v4
sdk: android-client-sdk
kind: scaffold
lang: kotlin
file: app/src/main/java/com/launchdarkly/hello_android/SnippetV4.kt
description: |
  Parse-and-type-check validator for Android Kotlin doc fragments that
  target the v4.x API surface (`LDConfig.Builder()` with no arguments —
  v5 made the `AutoEnvAttributes` constructor argument mandatory, so
  v4-era fragments cannot compile against the v5 aar the
  `android-client` validator container pre-bakes).

  Kotlin sibling of `java-syntax-only-v4-android`: stays inside the
  `android-client` container in `SNIPPET_CHECK=parse` mode and declares
  nested stub classes for just the v4 surface the fragments reference —
  the config builder plus the event config for private attributes, and
  the `LDClient.init` / `setOffline` surface and ambient `application` /
  `context` the offline-mode fragments use. The stubs are nested in a
  host class so the unqualified names resolve from the wrappee body
  without polluting the package namespace: other files in the pre-baked
  project reference the real v5 `LDConfig` / `LDClient`, and the file
  deliberately does NOT import `com.launchdarkly.sdk.android`, so the v5
  aar's same-named types never collide with the stubs. The shared
  `com.launchdarkly.sdk` types (`LDContext`) are real and version-stable
  across v4/v5. Stubs for `LDClient` (v4 init surface),
  `Components` / `EventProcessorBuilder` (v4 events config), and ambient
  `application` / `context` cover the multienv and private-attrs
  fragments that reference these names. The existing android sdk-docs CI
  row picks these snippets up with no workflow changes.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: android-client
  entrypoint: app/src/main/java/com/launchdarkly/hello_android/SnippetV4.kt
  env:
    SNIPPET_CHECK: parse
---

```kotlin
package com.launchdarkly.hello_android

import android.app.Application
import com.launchdarkly.sdk.LDContext
import java.util.concurrent.Future

// Stub of the v4-era android config + client surface. Only the members
// the doc fragments call are declared; everything returns a stub so the
// chained-builder shape type-checks. Nested inside the host class so the
// unqualified names resolve from the wrappee body without polluting the
// package namespace (other files in the pre-baked project reference the
// real v5 `LDConfig` / `LDClient`).
@Suppress("UNUSED_VARIABLE", "UNREACHABLE_CODE", "unused", "UNUSED_PARAMETER")
class SnippetV4 {
    class LDConfig {
        class Builder {
            fun mobileKey(key: String): Builder = this
            fun secondaryMobileKeys(keys: Map<String, String>): Builder = this
            fun offline(offline: Boolean): Builder = this
            fun events(eventsConfig: EventProcessorBuilder): Builder = this
            fun build(): LDConfig = LDConfig()
        }
    }

    class EventProcessorBuilder {
        fun allAttributesPrivate(allAttributesPrivate: Boolean): EventProcessorBuilder = this
        fun privateAttributes(vararg attributeNames: String): EventProcessorBuilder = this
    }

    object Components {
        fun sendEvents(): EventProcessorBuilder = EventProcessorBuilder()
    }

    // Stub of the v4 android LDClient init surface. The application
    // parameter is typed Application. The non-blocking overload (no
    // startWaitSeconds) returns a Future, matching the real v4 signature.
    class LDClient private constructor() {
        fun setOffline() {}
        companion object {
            fun init(
                application: Application,
                config: LDConfig,
                context: LDContext,
                startWaitSeconds: Int
            ): LDClient = LDClient()
            // The non-blocking v4 init overload (no startWaitSeconds)
            // returns a Future, matching the real v4 signature.
            fun init(
                application: Application,
                config: LDConfig,
                context: LDContext
            ): Future<LDClient> = TODO()
        }
    }

    // Ambient names the doc fragments assume an Activity host provides.
    // Never read at runtime (the body below is unreachable).
    @Suppress("UNUSED")
    val application: Application get() = TODO()
    @Suppress("UNUSED")
    val context: LDContext get() = TODO()

    fun wrappee() {
        if (false) {
{{ body }}
        }
    }
}
```
