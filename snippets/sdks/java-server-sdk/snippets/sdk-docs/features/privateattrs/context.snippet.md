---
id: java-server-sdk/sdk-docs/features/privateattrs/context
sdk: java-server-sdk
kind: reference
lang: java
description: Marking context attributes private with the context builder for Java SDK v6.0.
validation:
  scaffold: java-server-sdk/scaffolds/java-syntax-only

---

```java
LDContext context = LDContext.builder("example-context-key")
  .set("email", "sandy@example.com")
  .privateAttributes("email")
  .build();
```
