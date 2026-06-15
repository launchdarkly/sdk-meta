---
id: android-client-sdk/experimentation/track-only
sdk: android-client-sdk
kind: reference
lang: kotlin
description: Experimentation onboarding (track only) for android-client-sdk — initialize and add a trackMetric helper for conversion events.
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-syntax-only
---

```kotlin
import com.launchdarkly.sdk.*
import com.launchdarkly.sdk.android.*
import com.launchdarkly.sdk.android.LDConfig.Builder.AutoEnvAttributes

// This is your mobile key.
val ldConfig = LDConfig.Builder(AutoEnvAttributes.Enabled)
    .mobileKey("YOUR_MOBILE_KEY")
    .build()

// A "context" is a data object representing users, devices, organizations, and other entities.
val context = LDContext.create("EXAMPLE_CONTEXT_KEY")

// If you don't want to block execution while the SDK tries to get
// latest flags, move this code into an async IO task and await on its completion.
val client: LDClient = LDClient.init(this@BaseApplication, ldConfig, context, 5)

// Call trackMetric when a metric action occurs in your app —
// a tap, a form submit, a screen view, a custom event, whatever your metric measures.
fun trackMetric(metricKey: String, data: LDValue = LDValue.ofNull()) {
    LDClient.get().trackData(metricKey, data)
}
```
