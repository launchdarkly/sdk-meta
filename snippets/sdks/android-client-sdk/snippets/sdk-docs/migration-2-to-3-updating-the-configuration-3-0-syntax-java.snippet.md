---
id: android-client-sdk/sdk-docs/migration-2-to-3-updating-the-configuration-3-0-syntax-java
sdk: android-client-sdk
kind: reference
lang: java
description: "3.0 syntax (Java) in section \"Updating the configuration\""
---

```java
LDConfig ldConfig = new LDConfig.Builder()
    .mobileKey("example-mobile-key")
    .privateAttributes(
        UserAttribute.NAME,
        UserAttribute.forName("group")
    )
    .pollUri(Uri.parse("https://base.custom_domain.com"))
    .eventsUri(Uri.parse("https://events.custom_domain.com"))
    .headerTransform(headers -> {
        headers.put("Proxy-Authorization", "OTHER_AUTH")
    })
    .build();
```
