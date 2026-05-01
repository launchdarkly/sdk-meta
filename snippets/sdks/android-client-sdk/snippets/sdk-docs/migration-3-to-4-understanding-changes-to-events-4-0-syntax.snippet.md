---
id: android-client-sdk/sdk-docs/migration-3-to-4-understanding-changes-to-events-4-0-syntax
sdk: android-client-sdk
kind: reference
lang: java
description: "4.0 syntax in section \"Understanding changes to events\""
---

```java
// Customizing event behavior
LDConfig config = new LDConfig.Builder()
  .events(
    Components.sendEvents()
      .capacity(20000)
      .flushIntervalMillis(10000)
      .privateAttributes("email", "name", "myCustomAttribute")
  )
  .build();

// Disabling events
LDConfig config = new LDConfig.Builder()
  .events(Components.noEvents())
  .build();
```
