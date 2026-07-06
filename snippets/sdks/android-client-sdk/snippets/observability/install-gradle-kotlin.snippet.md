---
id: android-client-sdk/observability/install-gradle-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
file: android-client-sdk/observability/install-gradle-kotlin.gradle.kts
description: "Android observability plugin in section \"Install the plugin\" (Gradle Kotlin)"
# Not validated: Android AAR coordinates cannot be resolved by the shell-install plain-JVM gradle project; the android-client validator compiles against these exact artifacts.
---

```kotlin
implementation("com.launchdarkly:launchdarkly-android-client-sdk:5.+")
implementation("com.launchdarkly:launchdarkly-observability-android:0.60.0")
```
