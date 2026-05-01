---
id: android-client-sdk/sdk-docs/migration-2-to-3-track-methods-3-0-syntax-java
sdk: android-client-sdk
kind: reference
lang: java
description: "3.0 syntax (Java) in section \"Track methods\""
---

```java
client.trackData("dataEvent", LDValue.of(4));
client.trackMetric("metricEvent", LDValue.ofNull(), 5.5);
client.trackMetric("bothEvent", LDValue.of("tag"), 3.5);
```
