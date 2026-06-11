---
id: flutter-client-sdk/sdk-docs/features/evaluation-reasons/evaluation-reasons-v3
sdk: flutter-client-sdk
kind: reference
lang: dart
description: Flag evaluation reason example for Flutter SDK v3.x.
validation:
  scaffold: flutter-client-sdk/scaffolds/flutter-syntax-only-v3

---

```dart
LDConfig config = LDConfigBuilder('example-mobile-key', AutoEnvAttributes.Enabled)
    .evaluationReasons(true)
    .build();

// initialize client and context

LDEvaluationDetail<bool> detail =
  await LDClient.boolVariationDetail('example-flag-key', false);
  // or stringVariationDetail for a string-valued flag, and so on.

bool value = detail.value;
int index = detail.variationIndex;
LDEvaluationReason reason = detail.reason;
```
