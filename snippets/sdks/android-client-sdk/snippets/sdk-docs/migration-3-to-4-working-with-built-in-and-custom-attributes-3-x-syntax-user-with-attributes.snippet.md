---
id: android-client-sdk/sdk-docs/migration-3-to-4-working-with-built-in-and-custom-attributes-3-x-syntax-user-with-attributes
sdk: android-client-sdk
kind: reference
lang: java
description: "3.x syntax, user with attributes in section \"Working with built-in and custom attributes\""
---

```java
LDUser user = new LDUser.Builder("example-user-key")
  .name("Sandy")
  .email("sandy@example.com")
  .custom("group", "Global Health Services")
  .build();
```
