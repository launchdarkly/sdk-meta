---
id: apex-server-sdk/sdk-docs/features/privateattrs/config
sdk: apex-server-sdk
kind: reference
lang: java
description: Private attribute configuration for Apex.
validation:
  scaffold: apex-server-sdk/scaffolds/apex-syntax-only

---

```java
LDConfig config = new LDConfig.Builder()
    .setAllAttributesPrivate(true)
    .build();
```
