---
id: java-server-sdk/sdk-docs/migration-5-to-6-referencing-properties-of-an-attribute-object-6-0-syntax-context-with-object-attributes
sdk: java-server-sdk
kind: reference
lang: java
description: "6.0 syntax, context with object attributes in section \"Referencing properties of an attribute object\""
---

```java
LDValue addressData = LDValue.buildObject()
  .put("street", "Main St")
  .put("city", "Springfield")
  .build();
LDContext context = LDContext.builder("example-user-key")
  .set("address", addressData)
  .build();
```
