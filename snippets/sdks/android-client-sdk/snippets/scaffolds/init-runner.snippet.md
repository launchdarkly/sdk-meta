---
id: android-client-sdk/scaffolds/init-runner
sdk: android-client-sdk
kind: scaffold
lang: kotlin
file: app/src/main/java/com/launchdarkly/hello_android/MainApplication.kt
description: |
  Runs an `init.txt`-style Android Kotlin snippet end-to-end against
  a real LaunchDarkly mobile env. The body uses
  `this@BaseApplication` to hand an Application context to
  `LDClient.init(...)`, and that reference only resolves inside the
  body of an Application subclass. We splice the body into
  `onCreate()` of a `MainApplication: Application()` class — the
  same name the existing `android-client` validator harness's
  Robolectric test expects (`@Config(application =
  MainApplication::class)`).

  The validator's HelloAppTest then drives MainActivity through its
  full Robolectric lifecycle, observes the test flag, and asserts
  the canonical `feature flag evaluates to true` line lands in the
  TextView. Robolectric runs the Android framework inside the JVM
  — no emulator required, the existing docker harness handles
  everything.

  Layout:
    - MainApplication.kt (this scaffold) — splices the snippet body
      into onCreate. Substitutes the snippet's literal
      `YOUR_MOBILE_KEY` for the env-injected mobile key via
      `validation.placeholders`. Renames the body's
      `this@BaseApplication` to `this@MainApplication` so the
      Application reference resolves to the existing class name
      this scaffold produces. (Keeping the snippet body unchanged
      under the docs surface is essential; the rename happens only
      at validate time.)
    - MainActivity.kt (companion) — reads the flag via
      `LDClient.get()` once init finishes, observes for changes,
      and writes the canonical EXAM-HELLO success line into the
      TextView. The activity's layout (`R.layout.activity_main`,
      `R.id.textview`) comes from the upstream hello-android
      scaffold pre-baked into the validator image.
inputs:
  body:
    type: string
    description: The wrappee init snippet body, embedded after key + class-name substitution.
validation:
  runtime: android-client
  entrypoint: app/src/main/java/com/launchdarkly/hello_android/MainApplication.kt
  companions:
    - android-client-sdk/scaffolds/init-runner-activity
---

```kotlin
package com.launchdarkly.hello_android

import android.app.Application

class MainApplication : Application() {
    override fun onCreate() {
        super.onCreate()
{{ body }}
    }
}
```
