---
id: android-client-sdk/sdk-docs/migration-3-to-4-understanding-changes-to-private-attributes-4-0-syntax-all-attributes-private
sdk: android-client-sdk
kind: reference
lang: java
description: "4.0 syntax, all attributes private in section \"Understanding changes to private attributes\""
---

```java
LDConfig config = new LDConfig.Builder()
.events(
  Components.sendEvents().allAttributesPrivate(true)
)
.build();
```
