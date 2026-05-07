---
id: java-server-sdk/sdk-docs/evaluate-a-context-java-sdk-v6-0
sdk: java-server-sdk
kind: reference
lang: java
description: "Java SDK v6.0+ in section \"Evaluate a context\""
validation:
  scaffold: java-server-sdk/scaffolds/java-syntax-only
---

```java
LDContext context = LDContext.builder("example-context-key")
  .name("Sandy")
  .build();

boolean flagValue = client.boolVariation("example-flag-key", context, false);

if (flagValue) {
  // Application code to show the feature
}
else {
  // The code to run if the feature is off
}
```
