---
id: java-server-sdk/sdk-docs/features/anonymous/anonymous
sdk: java-server-sdk
kind: reference
lang: java
description: Anonymous context example for Java.
validation:
  scaffold: java-server-sdk/scaffolds/java-syntax-only
---

```java
LDContext context = LDContext.builder("example-context-key")
  .anonymous(true)
  .build();
```
