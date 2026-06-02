---
id: android-client-sdk/sdk-docs/initialize-the-client-android-sdk-v5-x-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
description: "Android SDK v5.x (Kotlin) in section \"Initialize the client\""
# TODO(snippet-bug): body calls `Observability(this@BaseApplication)`
# with just the Application context, but the v0.49.0
# `com.launchdarkly:launchdarkly-observability-android` constructor
# is `Observability(Application application, String mobileKey,
# ObservabilityOptions options)` — three required args. Confirmed by
# inspecting the AAR's
# `com/launchdarkly/observability/plugin/Observability.class`. The
# doc shape appears aspirational toward a future 1.x release where
# the constructor may be simplified. Fix in the snippet-bugs PR
# once the API stabilizes, or update the snippet to pass all three
# args for the current 0.x constructor.
---

```kotlin
val ldConfig = LDConfig.Builder(AutoEnvAttributes.Enabled)
    .mobileKey("example-mobile-key")
    // optional observability plugin, requires LaunchDarkly Android Client SDK v5.9+
    .plugins(Components.plugins().setPlugins(
      listOf(Observability(this@BaseApplication, "example-mobile-key"))
    ))
    // other options
    .build()

// You'll need this context later, but you can ignore it for now.
val context = LDContext.create("example-context-key")

val client: LDClient = LDClient.init(this@BaseApplication, ldConfig, context, 0)
```
