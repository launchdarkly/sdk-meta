---
id: java-server-sdk/sdk-docs/features/contextconfig/multi-context
sdk: java-server-sdk
kind: reference
lang: java
description: Multi-context example for Java SDK v6.0.
validation:
  scaffold: java-server-sdk/scaffolds/java-syntax-only

---

```java
LDContext multiContext = LDContext.createMulti(
    LDContext.create("example-user-key"),
    LDContext.create(ContextKind.of("device"), "example-device-key")
);
```
