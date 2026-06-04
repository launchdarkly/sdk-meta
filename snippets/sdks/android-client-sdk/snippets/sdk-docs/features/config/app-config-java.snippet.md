---
id: android-client-sdk/sdk-docs/features/config/app-config-java
sdk: android-client-sdk
kind: reference
lang: java
description: Application metadata configuration example for Android.
validation:
  scaffold: android-client-sdk/scaffolds/java-syntax-only

---

```java
LDConfig config = new LDConfig.Builder(AutoEnvAttributes.Enabled)
    .applicationInfo(
        Components.applicationInfo()
            .applicationId("authentication-service")
            .applicationName("Authentication-Service")
            .applicationVersion("1.0.0")
            .applicationVersionName("v1")
    )
    .build();
```
