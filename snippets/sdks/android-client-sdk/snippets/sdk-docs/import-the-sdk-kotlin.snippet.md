---
id: android-client-sdk/sdk-docs/import-the-sdk-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
description: "Kotlin in section \"Import the SDK\""
# Bucket C: jvm validator pulls launchdarkly-java-server-sdk, not the android-client SDK (which lives in Google Maven as an aar). See _sdk-docs-port-notes.md.
---

```kotlin
import com.launchdarkly.sdk.*
import com.launchdarkly.sdk.android.*

// optional observability plugin, requires LaunchDarkly Android Client SDK v5.9+
import com.launchdarkly.observability.plugin.Observability
import com.launchdarkly.sdk.android.integrations.Plugin
```
