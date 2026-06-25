---
id: apex-server-sdk/sdk-docs/features/anonymous/anonymous
sdk: apex-server-sdk
kind: reference
lang: java
description: Anonymous user example for Apex.
validation:
  scaffold: apex-server-sdk/scaffolds/apex-syntax-only
---

```java
LDUser user = new LDUser.Builder('example-user-key')
    .setAnonymous(true)
    .build();
```
