---
id: java-server-sdk/sdk-info/flagEval-string
sdk: java-server-sdk
kind: flag-eval
lang: java
file: java-server-sdk/flagEval-string.txt
description: Flag evaluation example for java-server-sdk (string flag).
validation:
  scaffold: java-server-sdk/scaffolds/java-syntax-only
---

```java
// Set up the evaluation context.
final LDContext context = LDContext.builder("example-context-key")
    .name("Sandy")
    .build();

// Evaluate the feature flag for this context.
String flagValue = client.stringVariation("featureKey", context, "fallback-value");

if (flagValue.equals("example-variation-value")) {

    // TODO: Put your feature here

} else {

    // TODO: Put your fallback behavior here

}
```
