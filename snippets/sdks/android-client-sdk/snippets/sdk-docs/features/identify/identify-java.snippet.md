---
id: android-client-sdk/sdk-docs/features/identify/identify-java
sdk: android-client-sdk
kind: reference
lang: java
description: Identify example for the Android SDK v4.0+ (Java).
validation:
  scaffold: android-client-sdk/scaffolds/java-syntax-only
---

```java
LDContext updatedContext = LDContext.builderFromContext(context)
    .set("email", "sandy@example.com")
    .build();

client.identify(updatedContext);
```
