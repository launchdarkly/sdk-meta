---
id: java-server-sdk/sdk-docs/openfeature/install-dependencies-gradle
sdk: java-server-sdk
kind: reference
lang: groovy
file: java-server-sdk/sdk-docs/openfeature/install-dependencies-gradle.gradle
description: "Gradle in section \"Install the provider and dependencies\" (SDK + OpenFeature)"
validation:
  runtime: shell-install
---

```groovy
implementation group: 'com.launchdarkly', name: 'launchdarkly-java-server-sdk', version: '[7.1.0, 8.0.0)'
implementation 'dev.openfeature:sdk:[1.7.0,2.0.0)'
```
