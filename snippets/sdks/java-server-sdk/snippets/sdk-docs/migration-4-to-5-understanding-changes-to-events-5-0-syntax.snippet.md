---
id: java-server-sdk/sdk-docs/migration-4-to-5-understanding-changes-to-events-5-0-syntax
sdk: java-server-sdk
kind: reference
lang: java
description: "5.0 syntax in section \"Understanding changes to events\""
---

```java
// 5.0 model: disabling events
LDConfig config = new LDConfig.Builder()
  .events(Components.noEvents())
  .build();

// 5.0 model: customizing event behavior
LDConfig config = new LDConfig.Builder()
  .events(
    Components.sendEvents()
      .capacity(20000)
      .flushInterval(Duration.ofSeconds(10))
      .privateAttributes(UserAttribute.EMAIL, UserAttribute.NAME,
          UserAttribute.forName("myCustomAttribute"))
  )
  .build();
```
