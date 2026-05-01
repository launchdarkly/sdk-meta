---
id: android-client-sdk/sdk-docs/migration-3-to-4-understanding-changes-to-alias-events-4-0-syntax-associating-two-contexts
sdk: android-client-sdk
kind: reference
lang: java
description: "4.0 syntax, associating two contexts in section \"Understanding changes to alias events\""
---

```java
LDContext context1 = LDContext.create("example-user-key");
LDContext context2 = LDContext.create(ContextKind.of("device"), "example-device-key");
LDContext multiContext = LDContext.createMulti(context1, context2);
client.identify(multiContext);
```
