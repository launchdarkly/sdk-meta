---
id: android-client-sdk/sdk-docs/features/contextconfig/multi-context-java
sdk: android-client-sdk
kind: reference
lang: java
description: Multi-context example for Android SDK v4.0+.
validation:
  scaffold: android-client-sdk/scaffolds/java-syntax-only

---

```java
LDContext userContext = LDContext.create("example-user-key");
LDContext deviceContext = LDContext.create(ContextKind.of("device"), "example-device-key");

LDContext multiContext = LDContext.createMulti(
    userContext,
    deviceContext
);
```
