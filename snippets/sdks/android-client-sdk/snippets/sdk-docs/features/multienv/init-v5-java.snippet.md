---
id: android-client-sdk/sdk-docs/features/multienv/init-v5-java
sdk: android-client-sdk
kind: reference
lang: java
description: Multi-environment configuration for Android SDK v5.x (Java).
validation:
  scaffold: android-client-sdk/scaffolds/java-syntax-only

---

```java
Map<String, String> otherKeys = new HashMap<String, String>();
otherKeys.put("platform", "platform-mobile-key-456def");
otherKeys.put("core", "core-mobile-key-789ghi");

LDConfig ldConfig = new LDConfig.Builder(AutoEnvAttributes.Enabled)
    .mobileKey("example-mobile-key")
    .secondaryMobileKeys(otherKeys)
    .build();

LDContext context = LDContext.builder("example-context-key")
    .set("email", "sandy@example.com")
    .build();

LDClient.init(this.getApplication(), ldConfig, context);
```
