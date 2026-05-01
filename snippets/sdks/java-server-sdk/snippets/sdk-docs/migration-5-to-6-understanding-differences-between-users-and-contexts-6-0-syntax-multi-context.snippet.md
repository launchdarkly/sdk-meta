---
id: java-server-sdk/sdk-docs/migration-5-to-6-understanding-differences-between-users-and-contexts-6-0-syntax-multi-context
sdk: java-server-sdk
kind: reference
lang: java
description: "6.0 syntax, multi-context in section \"Understanding differences between users and contexts\""
---

```java

LDContext userContext = LDContext.create("example-user-key");
LDContext deviceContext = LDContext.create(ContextKind.of("device"), "example-device-key");

LDContext multiContext = LDContext.createMulti(
    userContext,
    deviceContext
);
```
