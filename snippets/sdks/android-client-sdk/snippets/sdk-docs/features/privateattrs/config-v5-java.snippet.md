---
id: android-client-sdk/sdk-docs/features/privateattrs/config-v5-java
sdk: android-client-sdk
kind: reference
lang: java
description: Private attribute configuration for Android SDK v5.x (Java).
validation:
  scaffold: android-client-sdk/scaffolds/java-syntax-only

---

```java
// All attributes marked private
LDConfig ldConfig = new LDConfig.Builder(AutoEnvAttributes.Enabled)
    .mobileKey("example-mobile-key")
    .events(
      Components.sendEvents()
          .allAttributesPrivate(true)
    )
    .build();

// Two attributes marked private
ldConfig = new LDConfig.Builder(AutoEnvAttributes.Enabled)
    .mobileKey("example-mobile-key")
    .events(
        Components.sendEvents()
            .privateAttributes("name", "group")
    )
    .build();
```
