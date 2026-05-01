---
id: java-server-sdk/sdk-docs/migration-4-to-5-understanding-changes-to-events-4-x-syntax
sdk: java-server-sdk
kind: reference
lang: java
description: "4.x syntax in section \"Understanding changes to events\""
---

```java
// 4.x model: disabling events
LDConfig config = new LDConfig.Builder()
  .sendEvents(false)
  .build();

// 4.x model: customizing event behavior
LDConfig config = new LDConfig.Builder()
  .capacity(20000)
  .flushInterval(10)
  .privateAttributes("email", "name", "myCustomAttribute")
  .build();
```
