---
id: android-client-sdk/sdk-docs/features/privateattrs/context-java
sdk: android-client-sdk
kind: reference
lang: java
description: Marking context attributes private with the context builder for Android SDK v4.0+ (Java).
validation:
  scaffold: android-client-sdk/scaffolds/java-syntax-only

---

```java
LDContext context = LDContext.builder("example-context-key")
    .set("email", "sandy@example.com")
    .set("name", "Sandy")
    .set("group", "Global Health Services")
    .privateAttributes("name", "group")
    .build();
```
