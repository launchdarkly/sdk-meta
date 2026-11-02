---
id: java-server-sdk/sdk-docs/features/overrides/override-affected
sdk: java-server-sdk
kind: reference
lang: java
description: Check whether a flag override affected an evaluation in Java.
validation:
  scaffold: java-server-sdk/scaffolds/java-syntax-only
---

```java
EvaluationDetail<Boolean> detail =
    client.boolVariationDetail("my-boolean-flag-key", context, false);
if (detail.getReason().isOverrideAffected()) {
    System.out.println("An override affected this evaluation");
}
```
