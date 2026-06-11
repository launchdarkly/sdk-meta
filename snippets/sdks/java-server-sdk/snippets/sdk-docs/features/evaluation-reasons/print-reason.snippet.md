---
id: java-server-sdk/sdk-docs/features/evaluation-reasons/print-reason
sdk: java-server-sdk
kind: reference
lang: java
description: Reason-object inspection example for Java.
validation:
  scaffold: java-server-sdk/scaffolds/java-syntax-only-members

---

```java
void printReason(EvaluationReason reason) {
  switch (reason.getKind()) {
    case OFF:
      System.out.println("it's off");
      break;
    case FALLTHROUGH:
      System.out.println("fell through");
      break;
    case TARGET_MATCH:
      System.out.println("targeted");
      break;
    case RULE_MATCH:
      System.out.println("matched rule " + reason.getRuleIndex()
        + "/" + reason.getRuleId());
      break;
    case PREREQUISITE_FAILED:
      System.out.println("prereq failed: " + reason.getPrerequisiteKey());
      break;
    case ERROR:
      System.out.println("error: " + reason.getErrorKind());
  }
  // or, if all you want is a simple descriptive string:
  System.out.println(reason.toString());
}
```
