---
id: android-client-sdk/sdk-docs/features/logging/logging-java
sdk: android-client-sdk
kind: reference
lang: java
description: Timber debug logging example for Android (Java).
validation:
  scaffold: android-client-sdk/scaffolds/java-syntax-only

---

```java
if (BuildConfig.DEBUG) {
    Timber.plant(new Timber.DebugTree());
}
```
