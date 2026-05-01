---
id: android-client-sdk/sdk-docs/migration-2-to-3-updating-the-configuration-2-x-syntax-java
sdk: android-client-sdk
kind: reference
lang: java
description: "2.x syntax (Java) in section \"Updating the configuration\""
---

```java
HashMap<String, String> addHeaders = new HashMap<>();
addHeaders.put("Proxy-Authorization", "OTHER_AUTH");

HashSet<String> privateAttributes = new HashSet<>();
privateAttributes.add("name");
privateAttributes.add("group");

LDConfig ldConfig = new LDConfig.Builder()
    .setMobileKey("example-mobile-key")
    .setPrivateAttributeNames(privateAttributes)
    .setBaseUri(Uri.parse("https://base.custom_domain.com"))
    .setEventsUri(Uri.parse("https://events.custom_domain.com/mobile"))
    .setAdditionalHeaders(addHeaders)
    .build();
```
