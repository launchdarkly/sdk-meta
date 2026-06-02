---
id: android-client-sdk/sdk-docs/import-the-sdk-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
description: "Kotlin in section \"Import the SDK\""
# TODO(validator): body imports
# `com.launchdarkly.observability.plugin.Observability` and
# `com.launchdarkly.sdk.android.integrations.Plugin`. The
# observability AAR isn't on the android-client validator's
# classpath (pre-baked app/build.gradle only pulls
# launchdarkly-android-client-sdk + AndroidX), so the imports
# don't resolve. Fix by extending the validator's build.gradle to
# include `com.launchdarkly:launchdarkly-observability-android` —
# can land in a follow-up validator-update PR.
---

```kotlin
import com.launchdarkly.sdk.*
import com.launchdarkly.sdk.android.*

// optional observability plugin, requires LaunchDarkly Android Client SDK v5.9+
import com.launchdarkly.observability.plugin.Observability
import com.launchdarkly.sdk.android.integrations.Plugin
```
