---
id: apex-server-sdk/sdk-docs/features/evaluation-reasons/evaluation-reasons
sdk: apex-server-sdk
kind: reference
lang: java
description: Flag evaluation reason example for Apex.
validation:
  scaffold: apex-server-sdk/scaffolds/apex-syntax-only

---

```java
LDClient.EvaluationDetail details = new LDClient.EvaluationDetail();

Boolean value = client.boolVariation(user, 'your.feature.key', false, details);

/* inspect details here */
if (details.getReason().getKind() == EvaluationReason.Kind.OFF) {
    /* ... */
}
```
