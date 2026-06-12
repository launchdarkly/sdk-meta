---
id: android-client-sdk/sdk-docs/features/datasaving/lifecycle-switching-disabled-java
sdk: android-client-sdk
kind: reference
lang: java
description: Disable lifecycle-driven mode switching while keeping network-driven switching for Android (Java).
validation:
  scaffold: android-client-sdk/scaffolds/java-syntax-only
---

```java
LDConfig config = new LDConfig.Builder(AutoEnvAttributes.Enabled)
    .mobileKey("example-mobile-key")
    .dataSystem(
        Components.dataSystem()
            .automaticModeSwitching(
                DataSystemComponents.automaticModeSwitching()
                    .lifecycle(false)
                    .network(true)
                    .build()))
    .build();
```
