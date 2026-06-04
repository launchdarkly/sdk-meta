---
id: android-client-sdk/sdk-docs/features/config/service-endpoint-configuration-java-eu
sdk: android-client-sdk
kind: reference
lang: java
description: Service endpoint configuration example for Android.

---

```java
LDConfig ldConfig = new LDConfig.Builder(AutoEnvAttributes.Enabled)
  .mobileKey("example-mobile-key")
  .serviceEndpoints(
    Components.serviceEndpoints()
      .streaming("https://clientstream.eu.launchdarkly.com")
      .polling("https://clientsdk.eu.launchdarkly.com")
      .events("https://events.eu.launchdarkly.com")
  )
  .build();
```
