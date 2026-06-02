---
id: android-client-sdk/sdk-docs/import-the-sdk-java
sdk: android-client-sdk
kind: reference
lang: java
description: "Java in section \"Import the SDK\""
# TODO(snippet-bug): the last two `import` lines are missing trailing
# semicolons (Kotlin-style imports mistakenly pasted into a Java
# code block in the source MDX). javac rejects them at parse. Also
# the observability AAR isn't on the android-client validator's
# classpath, so even with semicolons the names wouldn't resolve.
# Fix in the snippet-bugs PR: add semicolons and add the
# launchdarkly-observability-android dependency to the validator's
# pre-baked app/build.gradle, or split into separate Java vs Kotlin
# snippets.
---

```java
import com.launchdarkly.sdk.*;
import com.launchdarkly.sdk.android.*;

// optional observability plugin, requires LaunchDarkly Android Client SDK v5.9+
import com.launchdarkly.observability.plugin.Observability
import com.launchdarkly.sdk.android.integrations.Plugin
```
