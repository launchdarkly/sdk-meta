---
id: android-client-sdk/sdk-docs/migration-2-to-3-updating-the-configuration-3-0-syntax-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
description: "3.0 syntax (Kotlin) in section \"Updating the configuration\""
---

```kotlin
val ldConfig: LDConfig = LDConfig.Builder()
    .mobileKey("example-mobile-key")
    .privateAttributes(
        UserAttribute.NAME,
        UserAttribute.forName("group")
    )
    .pollUri(Uri.parse("https://base.custom_domain.com"))
    .eventsUri(Uri.parse("https://events.custom_domain.com"))
    .headerTransform({ it.put("Proxy-Authorization", "OTHER_AUTH") })
    .build()
```
