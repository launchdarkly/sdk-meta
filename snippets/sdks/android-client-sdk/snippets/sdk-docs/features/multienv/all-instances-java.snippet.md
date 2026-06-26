---
id: android-client-sdk/sdk-docs/features/multienv/all-instances-java
sdk: android-client-sdk
kind: reference
lang: java
description: Calls that affect all LDClient instances for Android SDK v4.0+ (Java).
validation:
  scaffold: android-client-sdk/scaffolds/java-syntax-only

---

```java
LDClient coreInstance = LDClient.getForMobileKey("core");

// Calls affect all LDClient Instances
coreInstance.identify(context);
coreInstance.flush();
coreInstance.setOffline();
coreInstance.setOnline();
coreInstance.close();
```
