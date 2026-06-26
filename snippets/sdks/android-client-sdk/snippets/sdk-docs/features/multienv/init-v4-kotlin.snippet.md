---
id: android-client-sdk/sdk-docs/features/multienv/init-v4-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
description: Multi-environment configuration for Android SDK v4.x (Kotlin).
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-syntax-only-v4

---

```kotlin
val otherKeys: MutableMap<String, String> = HashMap()
otherKeys.put("platform", "platform-example-mobile-key")
otherKeys.put("core", "core-example-mobile-key")

val ldConfig: LDConfig = LDConfig.Builder()
    .mobileKey("example-mobile-key")
    .secondaryMobileKeys(otherKeys)
    .build()

val context: LDContext = LDContext.builder("example-context-key")
    .set("email", "sandy@example.com")
    .build()

LDClient.init(application, ldConfig, context)
```
