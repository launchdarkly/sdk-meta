---
id: android-client-sdk/getting-started/build-gradle
sdk: android-client-sdk
kind: manifest-fragment
lang: gradle
description: Gradle dependency entry to drop into app/build.gradle.
inputs:
  version:
    type: string
    description: SDK version. Defaults to '5.0.0' in gonfalon as a fallback when the async fetch hasn't completed.
    runtime-default: "5.0.0"
ld-application:
  slot: build-gradle
---

Add the LaunchDarkly SDK as a dependency in the `app/build.gradle` file:

```gradle
dependencies {
  ...
  implementation("com.launchdarkly:launchdarkly-android-client-sdk:{{ version }}")
}
```
