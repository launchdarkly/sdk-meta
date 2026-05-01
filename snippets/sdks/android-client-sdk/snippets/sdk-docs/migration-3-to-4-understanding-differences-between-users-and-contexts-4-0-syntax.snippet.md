---
id: android-client-sdk/sdk-docs/migration-3-to-4-understanding-differences-between-users-and-contexts-4-0-syntax
sdk: android-client-sdk
kind: reference
lang: java
description: "4.0 syntax in section \"Understanding differences between users and contexts\""
---

```java
LDContext userContext = LDContext.create("example-user-key");
LDContext deviceContext = LDContext.create(ContextKind.of("device"), "example-device-key");

LDContext multiContext = LDContext.createMulti(
    userContext,
    deviceContext
);
```
