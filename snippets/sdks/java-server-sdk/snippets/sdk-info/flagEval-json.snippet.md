---
id: java-server-sdk/sdk-info/flagEval-json
sdk: java-server-sdk
kind: flag-eval
lang: java
file: java-server-sdk/flagEval-json.txt
description: Flag evaluation example for java-server-sdk (JSON flag).
validation:
  scaffold: java-server-sdk/scaffolds/java-syntax-only
---

```java
// Set up the evaluation context.
final LDContext context = LDContext.builder("example-context-key")
    .name("Sandy")
    .build();

// Evaluate the feature flag for this context.
LDValue flagValue = client.jsonValueVariation("featureKey", context, LDValue.ofNull());

// Use the flag value to configure your feature
// TODO: Put your feature here
```
