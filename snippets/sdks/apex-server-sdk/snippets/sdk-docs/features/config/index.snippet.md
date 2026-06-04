---
id: apex-server-sdk/sdk-docs/features/config/index
sdk: apex-server-sdk
kind: reference
lang: java
description: SDK configuration example for Apex.
validation:
  scaffold: apex-server-sdk/scaffolds/apex-syntax-only

---

```java
LDConfig config = new LDConfig.Builder()
    .setAllAttributesPrivate(true)
    .build();
```
