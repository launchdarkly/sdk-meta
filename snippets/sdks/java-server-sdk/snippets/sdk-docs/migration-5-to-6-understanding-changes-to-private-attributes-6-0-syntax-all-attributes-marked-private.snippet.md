---
id: java-server-sdk/sdk-docs/migration-5-to-6-understanding-changes-to-private-attributes-6-0-syntax-all-attributes-marked-private
sdk: java-server-sdk
kind: reference
lang: java
description: "6.0 syntax, all attributes marked private in section \"Understanding changes to private attributes\""
---

```java
LDConfig configWithAllAttributesPrivate = new LDConfig.Builder()
  .events(
    Components.sendEvents()
      .allAttributesPrivate(true)
  )
  .build();
```
