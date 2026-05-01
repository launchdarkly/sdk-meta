---
id: android-client-sdk/sdk-info/init
sdk: android-client-sdk
kind: init
lang: kotlin
file: android-client-sdk/init.txt
description: Client initialization snippet for android-client-sdk.
---

```kotlin
import com.launchdarkly.sdk.*
import com.launchdarkly.sdk.android.*;

// This is your mobile key.
val ldConfig = LDConfig.Builder(AutoEnvAttributes.Enabled)
    .mobileKey("YOUR_MOBILE_KEY")
    .build()

// A "context" is a data object representing users, devices, organizations, and other entities.
val context = LDContext.create("EXAMPLE_CONTEXT_KEY")

// If you don't want to block execution while the SDK tries to get
// latest flags, move this code into an async IO task and await on its completion.
val client: LDClient = LDClient.init(this@BaseApplication, ldConfig, context, 5)
```
