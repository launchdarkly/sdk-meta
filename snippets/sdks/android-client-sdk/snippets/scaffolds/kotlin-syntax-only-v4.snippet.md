---
id: android-client-sdk/scaffolds/kotlin-syntax-only-v4
sdk: android-client-sdk
kind: scaffold
lang: kotlin
file: app/src/main/java/com/launchdarkly/hello_android/SnippetV4.kt
description: |
  Parse-and-type-check validator for Android Kotlin doc fragments that
  target the v4.x configuration surface (`LDConfig.Builder()` with no
  arguments — v5 made the `AutoEnvAttributes` constructor argument
  mandatory, so v4-era fragments cannot compile against the v5 aar the
  `android-client` validator container pre-bakes).

  Kotlin sibling of `java-syntax-only-v4-android`: stays inside the
  `android-client` container in `SNIPPET_CHECK=parse` mode and
  declares nested stub classes for just the v4 builder surface the
  fragments reference. The file deliberately does NOT import
  `com.launchdarkly.sdk.android`, so the v5 aar's same-named types
  never collide with the stubs. The existing android sdk-docs CI row
  picks these snippets up with no workflow changes.

  Shared `com.launchdarkly.sdk` types (`LDContext`) are real — they
  ship inside the v5 aar's java-sdk-common dependency and are
  version-stable across v4/v5. Stubs for `LDClient` (v4 init surface),
  `Components` / `EventProcessorBuilder` (v4 events config), and ambient
  `application` / `context` cover the multienv and private-attrs
  fragments that reference these names.
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

import com.launchdarkly.sdk.LDContext
import java.util.concurrent.Future

// Stub of the v4-era android config surface. Only the members the doc
// fragments call are declared; everything returns a stub so the
// chained-builder shape type-checks. Nested inside the host class so
// the unqualified names resolve from the wrappee body without
// polluting the package namespace (other files in the pre-baked
// project reference the real v5 `LDConfig`).
@Suppress("UNUSED_VARIABLE", "UNREACHABLE_CODE", "unused", "UNUSED_PARAMETER")
class SnippetV4 {
    class LDConfig {
        class Builder {
            fun mobileKey(key: String): Builder = this
            fun secondaryMobileKeys(keys: Map<String, String>): Builder = this
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
    // parameter is typed Any so the ambient `application` stub below
    // satisfies it. The non-blocking overload (no startWaitSeconds)
    // returns a Future, matching the real v4 signature.
    class LDClient private constructor() {
        companion object {
            fun init(
                application: Any,
                config: LDConfig,
                context: LDContext,
                startWaitSeconds: Int
            ): LDClient = LDClient()
            // The non-blocking v4 init overload (no startWaitSeconds)
            // returns a Future, matching the real v4 signature.
            fun init(
                application: Any,
                config: LDConfig,
                context: LDContext
            ): Future<LDClient> = TODO()
        }
    }

    // Ambient names the doc fragments assume an Activity host provides.
    // Never read at runtime (the body below is unreachable).
    val application: Any get() = TODO()
    val context: LDContext get() = TODO()

    fun wrappee() {
        if (false) {
{{ body }}
        }
    }
}
```
