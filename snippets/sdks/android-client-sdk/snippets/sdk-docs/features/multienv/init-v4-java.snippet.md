---
id: android-client-sdk/sdk-docs/features/multienv/init-v4-java
sdk: android-client-sdk
kind: reference
lang: java
description: Multi-environment configuration for Android SDK v4.x (Java).
validation:
  scaffold: android-client-sdk/scaffolds/java-syntax-only-v4-android

---

```java
Map<String, String> otherKeys = new HashMap<String, String>();
otherKeys.put("platform", "platform-mobile-key-456def");
otherKeys.put("core", "core-mobile-key-789ghi");

LDConfig ldConfig = new LDConfig.Builder()
    .mobileKey("example-mobile-key")
    .secondaryMobileKeys(otherKeys)
    .build();

LDContext context = LDContext.builder("example-context-key")
    .set("email", "sandy@example.com")
    .build();

LDClient.init(this.getApplication(), ldConfig, context);
```
