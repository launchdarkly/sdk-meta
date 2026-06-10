---
id: flutter-client-sdk/sdk-docs/features/evaluation-reasons/evaluation-reasons-v4
sdk: flutter-client-sdk
kind: reference
lang: dart
description: Flag evaluation reason example for Flutter SDK v4.
validation:
  scaffold: flutter-client-sdk/scaffolds/flutter-syntax-only

---

```dart
final config = LDConfig(
  CredentialSource.fromEnvironment(),
  AutoEnvAttributes.enabled,
  dataSourceConfig: DataSourceConfig(
    evaluationReasons: true
  ),
);

// initialize client and context

LDEvaluationDetail<bool> detail =
  client.boolVariationDetail('example-flag-key', false);
  // or stringVariationDetail for a string-valued flag, and so on.

bool value = detail.value;
int? index = detail.variationIndex;
LDEvaluationReason? reason = detail.reason;
```
