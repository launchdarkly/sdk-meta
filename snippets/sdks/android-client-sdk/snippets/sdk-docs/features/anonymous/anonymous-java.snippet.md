---
id: android-client-sdk/sdk-docs/features/anonymous/anonymous-java
sdk: android-client-sdk
kind: reference
lang: java
description: Anonymous context example for Android (Java).
validation:
  scaffold: android-client-sdk/scaffolds/java-syntax-only
---

```java
LDContext context = LDContext.builder("example-context-key")
    .anonymous(true)
    .build();
```
