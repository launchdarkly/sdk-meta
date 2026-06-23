---
id: android-client-sdk/sdk-docs/features/contextconfig/context-example-java
sdk: android-client-sdk
kind: reference
lang: java
description: Context example for Android SDK v4.0+ (Java).
validation:
  scaffold: android-client-sdk/scaffolds/java-syntax-only

---

```java
LDContext context = LDContext.builder("example-context-key")
    .set("email", "sandy@example.com")
    .set("firstName", "Sandy")
    .set("lastName", "Smith")
    .set("group", "Global Health Services")
    .build();
```
