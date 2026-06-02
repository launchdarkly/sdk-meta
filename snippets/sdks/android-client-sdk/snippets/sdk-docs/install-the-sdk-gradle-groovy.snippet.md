---
id: android-client-sdk/sdk-docs/install-the-sdk-gradle-groovy
sdk: android-client-sdk
kind: reference
lang: java
description: "Gradle Groovy in section \"Install the SDK\""
# TODO(snippet-bug): body is Gradle Groovy DSL
# (`implementation '…' ` calls into a Gradle Project's dependencies
# block), not Java. javac rejects at parse — no statements outside
# methods or classes. Mistagged as `java` in the source MDX. Fix in
# the snippet-bugs PR: re-tag as `groovy` (or a custom
# `gradle.groovy`) and route through a Gradle/Groovy parse path,
# or skip syntax validation.
---

```java
implementation 'com.launchdarkly:launchdarkly-android-client-sdk:5.+'

// optional observability plugin, requires LaunchDarkly Android Client SDK v5.9+
implementation 'com.launchdarkly:launchdarkly-observability-android:0.5.0'
```
