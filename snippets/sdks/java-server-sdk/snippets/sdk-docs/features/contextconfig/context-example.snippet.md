---
id: java-server-sdk/sdk-docs/features/contextconfig/context-example
sdk: java-server-sdk
kind: reference
lang: java
description: Context example for Java SDK v6.0.
validation:
  scaffold: java-server-sdk/scaffolds/java-syntax-only

---

```java
LDContext context = LDContext.builder("example-context-key")
  .set("firstName", "Sandy")
  .set("lastName", "Smith")
  .set("email", "sandy@example.com")
  .set("groups",
    LDValue.buildArray().add("Acme").add("Global Health Services").build())
  .build();
```
