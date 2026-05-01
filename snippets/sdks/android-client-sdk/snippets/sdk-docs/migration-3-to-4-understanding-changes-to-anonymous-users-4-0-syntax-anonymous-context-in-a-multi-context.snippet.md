---
id: android-client-sdk/sdk-docs/migration-3-to-4-understanding-changes-to-anonymous-users-4-0-syntax-anonymous-context-in-a-multi-context
sdk: android-client-sdk
kind: reference
lang: java
description: "4.0 syntax, anonymous context in a multi-context in section \"Understanding changes to anonymous users\""
---

```java
LDContext userContext = LDContext.builder("unknown-context-key")
    .anonymous(true)
    .build();
LDContext orgContext = LDContext.create(ContextKind.of("organization"), "example-organization-key");
LDContext multiContext = LDContext.createMulti(userContext, orgContext);
```
