---
id: java-server-sdk/sdk-info/flagEval
sdk: java-server-sdk
kind: flag-eval
lang: java
file: java-server-sdk/flagEval.txt
description: Flag evaluation example for java-server-sdk.
validation:
  scaffold: java-server-sdk/scaffolds/java-syntax-only
---

```java
// Set up the evaluation context.
final LDContext context = LDContext.builder("example-context-key")
    .name("Sandy")
    .build();

// Evaluate the feature flag for this context.
boolean flagValue = client.boolVariation("featureKey", context, false);

if (flagValue) {

    // TODO: Put your feature here

} else {

    // TODO: Put your fallback behavior here

}
```
