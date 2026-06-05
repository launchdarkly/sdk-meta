---
id: android-client-sdk/sdk-docs/features/config/index-java
sdk: android-client-sdk
kind: reference
lang: java
description: SDK configuration example for Android.
validation:
  scaffold: android-client-sdk/scaffolds/java-syntax-only

---

```java
LDConfig ldConfig = new LDConfig.Builder(AutoEnvAttributes.Enabled)
    .mobileKey("example-mobile-key")
    .http(
      Components.httpConfiguration()
      .connectTimeoutMillis(5000)
    )
    .events(
      Components.sendEvents()
      .flushIntervalMillis(5000)
    )
    .build();
```
