---
id: android-client-sdk/sdk-docs/features/evaluation-reasons/print-reason-java
sdk: android-client-sdk
kind: reference
lang: java
description: Reason-object inspection example for Android (Java).
validation:
  scaffold: android-client-sdk/scaffolds/java-syntax-only-members

---

```java
void printReason(EvaluationReason reason) {
  switch (reason.getKind()) {
    case OFF:
      Timber.d("it's off");
      break;
    case FALLTHROUGH:
      Timber.d("fell through");
      break;
    case TARGET_MATCH:
      Timber.d("targeted");
      break;
    case RULE_MATCH:
      Timber.d("matched rule %d/%s",
               reason.getRuleIndex(),
               reason.getRuleId());
      break;
    case PREREQUISITE_FAILED:
      Timber.d("prereq failed: %s", reason.getPrerequisiteKey());
      break;
    case ERROR:
      Timber.d("error: %s", reason.getErrorKind());
  }
  // or, if all you want is a simple descriptive string:
  Timber.d(reason.toString());
}
```
