---
id: java-server-sdk/sdk-docs/migration-5-to-6-understanding-changes-to-alias-events-6-0-syntax-associating-two-contexts
sdk: java-server-sdk
kind: reference
lang: java
description: "6.0 syntax, associating two contexts in section \"Understanding changes to alias events\""
---

```java
LDContext context1 = LDContext.create("example-user-key");
LDContext context2 = LDContext.create(ContextKind.of("device"), "example-device-key");
LDContext multiContext = LDContext.createMulti(context1, context2);
client.identify(multiContext);
```
