---
id: android-client-sdk/sdk-docs/migration-2-to-3-track-methods-2-x-syntax-java
sdk: android-client-sdk
kind: reference
lang: java
description: "2.x syntax (Java) in section \"Track methods\""
---

```java
client.track("dataEvent", new JsonPrimitive(4));
client.track("metricEvent", null, 5.5);
client.track("bothEvent", new JsonPrimitive("tag"), 3.5);
```
