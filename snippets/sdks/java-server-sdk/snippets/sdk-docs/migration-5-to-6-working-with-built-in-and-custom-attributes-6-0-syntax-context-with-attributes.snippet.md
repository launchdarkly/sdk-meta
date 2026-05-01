---
id: java-server-sdk/sdk-docs/migration-5-to-6-working-with-built-in-and-custom-attributes-6-0-syntax-context-with-attributes
sdk: java-server-sdk
kind: reference
lang: java
description: "6.0 syntax, context with attributes in section \"Working with built-in and custom attributes\""
---

```java
LDContext context = LDContext.builder("example-user-key")
  .name("Sandy")
  .set("email", "sandy@example.com")
  .set("groups",
    LDValue.buildArray().add("Acme").add("Global Health Services").build())
  .build();
```
