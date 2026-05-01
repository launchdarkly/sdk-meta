---
id: java-server-sdk/sdk-docs/migration-5-to-6-understanding-changes-to-private-attributes-6-0-syntax-attribute-marked-private-for-one-context
sdk: java-server-sdk
kind: reference
lang: java
description: "6.0 syntax, attribute marked private for one context in section \"Understanding changes to private attributes\""
---

```java
LDContext context = LDContext.builder("key")
  .name("Sandy")
  .set("email", "sandy@example.com")
  .privateAttributes("email")
  .build();
```
