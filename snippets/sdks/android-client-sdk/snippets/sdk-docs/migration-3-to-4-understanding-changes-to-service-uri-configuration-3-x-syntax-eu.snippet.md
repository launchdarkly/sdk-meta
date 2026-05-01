---
id: android-client-sdk/sdk-docs/migration-3-to-4-understanding-changes-to-service-uri-configuration-3-x-syntax-eu
sdk: android-client-sdk
kind: reference
lang: java
description: "3.x syntax, EU in section \"Understanding changes to service URI configuration\""
---

```java
LDConfig config = new LDConfig.Builder()
  .pollUri(Uri.parse("https://clientsdk.eu.launchdarkly.com"))
  .streamUri(Uri.parse("https://clientstream.eu.launchdarkly.com"))
  .eventsUri(Uri.parse("https://events.eu.launchdarkly.com"))
  .build();
```
