---
id: android-client-sdk/sdk-docs/features/datasaving/disable-mode-switching-java
sdk: android-client-sdk
kind: reference
lang: java
description: Disable automatic mode switching entirely for Android (Java).
validation:
  scaffold: android-client-sdk/scaffolds/java-syntax-only
---

```java
LDConfig config = new LDConfig.Builder(AutoEnvAttributes.Enabled)
    .mobileKey("example-mobile-key")
    .dataSystem(
        Components.dataSystem()
            .automaticModeSwitching(AutomaticModeSwitchingConfig.disabled())
            .foregroundConnectionMode(ConnectionMode.STREAMING))
    .build();
```
