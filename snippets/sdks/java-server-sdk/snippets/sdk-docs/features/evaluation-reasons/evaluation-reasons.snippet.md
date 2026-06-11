---
id: java-server-sdk/sdk-docs/features/evaluation-reasons/evaluation-reasons
sdk: java-server-sdk
kind: reference
lang: java
description: Flag evaluation reason example for Java.
validation:
  scaffold: java-server-sdk/scaffolds/java-syntax-only

---

```java
import com.launchdarkly.sdk.*;

EvaluationDetail<Boolean> detail =
  client.boolVariationDetail("example-flag-key", context, false);
  // or stringVariationDetail for a string-valued flag, and so on.

boolean value = detail.getValue();
int index = detail.getVariationIndex();       // will be < 0 if evaluation failed
EvaluationReason reason = detail.getReason();
```
