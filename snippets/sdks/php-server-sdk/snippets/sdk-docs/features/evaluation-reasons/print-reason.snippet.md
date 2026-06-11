---
id: php-server-sdk/sdk-docs/features/evaluation-reasons/print-reason
sdk: php-server-sdk
kind: reference
lang: php
description: Reason-object inspection example for PHP.
validation:
  scaffold: php-server-sdk/scaffolds/php-syntax-only

---

```php
function printReason($reason) {
  switch ($reason->getKind()) {
    case EvaluationReason::OFF:
      echo("it's off");
      break;
    case EvaluationReason::FALLTHROUGH:
      echo("fell through");
      break;
    case EvaluationReason::TARGET_MATCH:
      echo("targeted");
      break;
    case EvaluationReason::RULE_MATCH:
      echo("matched rule " . $reason->getRuleIndex() .
        "/" . $reason->getRuleId());
      break;
    case EvaluationReason::PREREQUISITE_FAILED:
      echo("prereq failed: " . $reason->getPrerequisiteKey());
      break;
    case EvaluationReason::ERROR:
      echo("error: " . $reason->getErrorKind());
      break;
  }
  // or, if all you want is a simple descriptive string:
  echo $reason;
}
```
