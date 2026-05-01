---
id: android-client-sdk/sdk-docs/migration-3-to-4-referencing-properties-of-an-attribute-object-4-0-syntax-context-with-object-attributes
sdk: android-client-sdk
kind: reference
lang: java
description: "4.0 syntax, context with object attributes in section \"Referencing properties of an attribute object\""
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
