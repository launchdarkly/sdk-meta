---
id: android-client-sdk/sdk-docs/features/localstorage/localstorage-java
sdk: android-client-sdk
kind: reference
lang: java
description: Local storage caching example for Android (Java).
validation:
  scaffold: android-client-sdk/scaffolds/java-syntax-only
---

```java
LDConfig ldConfig = new LDConfig.Builder(AutoEnvAttributes.Enabled)
    .mobileKey("example-mobile-key")
    // Local storage is enabled by default
    // You can optionally configure the maximum number of cached contexts (default is 5)
    .maxCachedContexts(3)
    .build();

LDContext context = LDContext.create("example-context-key");

LDClient client = LDClient.init(this.getApplication(), ldConfig, context, 0);
```
