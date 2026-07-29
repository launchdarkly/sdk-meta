---
id: java-server-sdk/sdk-info/flagEval-number
sdk: java-server-sdk
kind: flag-eval
lang: java
file: java-server-sdk/flagEval-number.txt
description: Flag evaluation example for java-server-sdk (number flag).
validation:
  scaffold: java-server-sdk/scaffolds/java-syntax-only
---

```java
// Set up the evaluation context.
final LDContext context = LDContext.builder("example-context-key")
    .name("Sandy")
    .build();

// Evaluate the feature flag for this context.
double flagValue = client.doubleVariation("featureKey", context, 0);

if (flagValue == 1) {

    // TODO: Put your feature here

} else {

    // TODO: Put your fallback behavior here

}
```
