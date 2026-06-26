---
id: android-client-sdk/sdk-docs/features/relay-proxy-config/proxy-mode/proxy-mode-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
description: Proxy mode configuration example for Android (Kotlin).
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-syntax-only
---

```kotlin
val ldConfig: LDConfig = LDConfig.Builder(AutoEnvAttributes.Enabled)
  .mobileKey("example-mobile-key")
  .serviceEndpoints(
    Components.serviceEndpoints()
      .relayProxy("https://your-relay-proxy.com:8030")
  )
  .build()
```
