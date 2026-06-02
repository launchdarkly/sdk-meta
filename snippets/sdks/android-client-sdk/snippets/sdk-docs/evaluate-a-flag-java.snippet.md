---
id: android-client-sdk/sdk-docs/evaluate-a-flag-java
sdk: android-client-sdk
kind: reference
lang: java
description: "Java in section \"Evaluate a flag\""
# TODO(validate): jvm validator pulls launchdarkly-java-server-sdk, not the android-client SDK (which lives in Google Maven as an aar). See _sdk-docs-port-notes.md.
---

```java
boolean showFeature = client.boolVariation(flagKey, true);
if (showFeature) {
    // Application code to show the feature
}
else {
    // The code to run if the feature is off
}
```
