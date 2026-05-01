---
id: android-client-sdk/sdk-docs/migration-3-to-4-understanding-changes-to-service-uri-configuration-3-x-syntax-federal
sdk: android-client-sdk
kind: reference
lang: java
description: "3.x syntax, federal in section \"Understanding changes to service URI configuration\""
---

```java
LDConfig config = new LDConfig.Builder()
  .pollUri(Uri.parse("https://clientsdk.launchdarkly.us"))
  .streamUri(Uri.parse("https://clientstream.launchdarkly.us"))
  .eventsUri(Uri.parse("https://events.launchdarkly.us"))
  .build();
```
