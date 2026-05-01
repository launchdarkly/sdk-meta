---
id: android-client-sdk/sdk-docs/evaluate-a-flag-java
sdk: android-client-sdk
kind: reference
lang: java
description: "Java in section \"Evaluate a flag\""
validation:
  scaffold: android-client-sdk/scaffolds/android-syntax-only
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
