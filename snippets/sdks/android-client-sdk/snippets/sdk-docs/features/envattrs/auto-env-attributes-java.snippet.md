---
id: android-client-sdk/sdk-docs/features/envattrs/auto-env-attributes-java
sdk: android-client-sdk
kind: reference
lang: java
description: Automatic environment attributes configuration for Android (Java).
validation:
  scaffold: android-client-sdk/scaffolds/java-syntax-only

---

```java
LDConfig ldConfig = new LDConfig.Builder(AutoEnvAttributes.Enabled)
    .mobileKey("example-mobile-key")
    .build();
```
