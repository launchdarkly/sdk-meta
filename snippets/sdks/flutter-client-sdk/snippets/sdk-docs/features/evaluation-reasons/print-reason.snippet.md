---
id: flutter-client-sdk/sdk-docs/features/evaluation-reasons/print-reason
sdk: flutter-client-sdk
kind: reference
lang: dart
description: Reason-object inspection example for Flutter (Dart).
validation:
  scaffold: flutter-client-sdk/scaffolds/flutter-syntax-only

---

```dart
void printReason(LDEvaluationReason reason) {
  switch (reason.kind) {
    case LDKind.off:
      print("it's off");
      break;
    case LDKind.fallthrough:
      print('fell through');
      break;
    case LDKind.targetMatch:
      print('targeted');
      break;
    case LDKind.ruleMatch:
      print('matched rule: ${reason.ruleIndex} ${reason.ruleId}');
      break;
    case LDKind.prerequisiteFailed:
      print('prereq failed: ${reason.prerequisiteKey}');
      break;
    case LDKind.error:
      print('error: ${reason.errorKind}');
      break;
  }
}
```
