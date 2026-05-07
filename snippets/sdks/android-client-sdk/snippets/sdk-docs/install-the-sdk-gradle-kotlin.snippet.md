---
id: android-client-sdk/sdk-docs/install-the-sdk-gradle-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
description: "Gradle Kotlin in section \"Install the SDK\""
# Bucket C: jvm validator pulls launchdarkly-java-server-sdk, not the android-client SDK (which lives in Google Maven as an aar). See _sdk-docs-port-notes.md.
---

```kotlin
implementation("com.launchdarkly:launchdarkly-android-client-sdk:5.+")

// optional observability plugin, requires LaunchDarkly Android Client SDK v5.9+
implementation("com.launchdarkly:launchdarkly-observability-android:0.5.0")
```
