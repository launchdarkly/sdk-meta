---
id: android-client-sdk/sdk-docs/install-the-sdk-gradle-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
description: "Gradle Kotlin in section \"Install the SDK\""
# TODO(snippet-bug): body is Gradle Kotlin DSL (`implementation("…")`
# calls into a build.gradle.kts dependencies block), not standalone
# Kotlin. The wrappee scope (`if (false) { ... }` inside
# `onCreate`) doesn't expose an `implementation()` method, so
# kotlinc reports unresolved reference. Mistagged as standalone
# `kotlin` in the source MDX. Fix in the snippet-bugs PR: re-tag
# (e.g. `gradle.kotlin`) and route through a build-script parse
# path, or skip syntax validation.
---

```kotlin
implementation("com.launchdarkly:launchdarkly-android-client-sdk:5.+")

// optional observability plugin, requires LaunchDarkly Android Client SDK v5.9+
implementation("com.launchdarkly:launchdarkly-observability-android:0.5.0")
```
