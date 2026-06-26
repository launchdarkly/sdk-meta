---
id: android-client-sdk/sdk-docs/features/multienv/get-for-mobile-key-java
sdk: android-client-sdk
kind: reference
lang: java
description: Accessing a secondary environment client instance on Android (Java).
validation:
  scaffold: android-client-sdk/scaffolds/java-syntax-only

---

```java
LDClient coreInstance = LDClient.getForMobileKey("core");
// Variation determines whether or not a flag is enabled for a specific context
coreInstance.boolVariation("core-flag", false);
// allFlags produces a map of feature flag keys to their values
coreInstance.allFlags();
// trackData records actions end users take in your app
coreInstance.trackData("example-metric-key", data);
```
