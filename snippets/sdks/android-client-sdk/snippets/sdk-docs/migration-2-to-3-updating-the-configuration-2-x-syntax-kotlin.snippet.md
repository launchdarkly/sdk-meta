---
id: android-client-sdk/sdk-docs/migration-2-to-3-updating-the-configuration-2-x-syntax-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
description: "2.x syntax (Kotlin) in section \"Updating the configuration\""
---

```kotlin
val addHeaders: HashMap<String, String> = HashMap()
addHeaders.put("Proxy-Authorization", "OTHER_AUTH")

val privateAttributes: HashSet<String> = new HashSet()
privateAttributes.add("name")
privateAttributes.add("group")

val ldConfig: LDConfig = LDConfig.Builder()
    .setMobileKey("example-mobile-key")
    .setPrivateAttributeNames(privateAttributes)
    .setBaseUri(Uri.parse("https://base.custom_domain.com"))
    .setEventsUri(Uri.parse("https://events.custom_domain.com/mobile"))
    .setAdditionalHeaders(addHeaders)
    .build()
```
