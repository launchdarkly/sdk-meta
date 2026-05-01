---
id: android-client-sdk/sdk-docs/migration-3-to-4-understanding-changes-to-private-attributes-4-0-syntax-attribute-marked-private-for-one-context
sdk: android-client-sdk
kind: reference
lang: java
description: "4.0 syntax, attribute marked private for one context in section \"Understanding changes to private attributes\""
---

```java
LDContext context = LDContext.builder("key")
  .name("Sandy")
  .set("email", "sandy@example.com")
  .privateAttributes("email")
  .build();
```
