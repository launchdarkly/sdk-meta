---
id: android-client-sdk/sdk-docs/migration-3-to-4-understanding-changes-to-anonymous-users-4-0-syntax-building-an-anonymous-context
sdk: android-client-sdk
kind: reference
lang: java
description: "4.0 syntax, building an anonymous context in section \"Understanding changes to anonymous users\""
---

```java
LDContext context = LDContext.builder("unknown-context-key")
    .anonymous(true)
    .build();
```
