---
id: java-server-sdk/sdk-docs/migration-5-to-6-working-with-built-in-and-custom-attributes-5-x-syntax-user-with-attributes
sdk: java-server-sdk
kind: reference
lang: java
description: "5.x syntax, user with attributes in section \"Working with built-in and custom attributes\""
---

```java
LDUser user = new LDUser.Builder("example-user-key")
  .name("Sandy")
  .email("sandy@example.com")
  .custom("groups",
    LDValue.buildArray().add("Acme").add("Global Health Services").build())
  .build();
```
