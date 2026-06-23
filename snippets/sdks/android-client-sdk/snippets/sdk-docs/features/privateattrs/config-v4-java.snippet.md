---
id: android-client-sdk/sdk-docs/features/privateattrs/config-v4-java
sdk: android-client-sdk
kind: reference
lang: java
description: Private attribute configuration for Android SDK v4.x (Java).
validation:
  scaffold: android-client-sdk/scaffolds/java-syntax-only-v4-android

---

```java
// All attributes marked private
LDConfig config = new LDConfig.Builder()
    .mobileKey("example-mobile-key")
    .events(
      Components.sendEvents()
          .allAttributesPrivate(true)
    )
    .build();

// Two attributes marked private
LDConfig ldConfig = new LDConfig.Builder()
    .mobileKey("example-mobile-key")
    .events(
        Components.sendEvents()
            .privateAttributes("name", "group")
    )
    .build();
```
