---
id: java-server-sdk/sdk-docs/features/privateattrs/config
sdk: java-server-sdk
kind: reference
lang: java
description: Private attribute configuration for Java SDK v6.0.
validation:
  scaffold: java-server-sdk/scaffolds/java-syntax-only

---

```java
// All attributes marked private
LDConfig configWithAllAttributesPrivate = new LDConfig.Builder()
  .events(
    Components.sendEvents()
      .allAttributesPrivate(true)
  )
  .build();

// Some attributes marked private
LDConfig configWithSpecificAttributesPrivate = new LDConfig.Builder()
  .events(
    Components.sendEvents()
      .privateAttributes("name", "email", "someAttribute")
  )
  .build();
```
