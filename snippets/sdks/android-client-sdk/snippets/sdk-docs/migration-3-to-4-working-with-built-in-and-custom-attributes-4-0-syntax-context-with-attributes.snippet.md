---
id: android-client-sdk/sdk-docs/migration-3-to-4-working-with-built-in-and-custom-attributes-4-0-syntax-context-with-attributes
sdk: android-client-sdk
kind: reference
lang: java
description: "4.0 syntax, context with attributes in section \"Working with built-in and custom attributes\""
---

```java
LDContext context = LDContext.builder("example-user-key")
  .name("Sandy")
  .set("email", "sandy@example.com")
  .set("group", "Global Health Services")
  .build();
```
