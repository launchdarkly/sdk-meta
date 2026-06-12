---
id: android-client-sdk/sdk-docs/features/datasaving/standard-setup-java
sdk: android-client-sdk
kind: reference
lang: java
description: Data saving mode standard setup for Android (Java).
validation:
  scaffold: android-client-sdk/scaffolds/java-syntax-only
---

```java
LDConfig config = new LDConfig.Builder(AutoEnvAttributes.Enabled)
    .mobileKey("example-mobile-key")
    .dataSystem(Components.dataSystem())
    .build();
```
