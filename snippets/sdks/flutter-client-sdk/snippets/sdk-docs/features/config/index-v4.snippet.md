---
id: flutter-client-sdk/sdk-docs/features/config/index-v4
sdk: flutter-client-sdk
kind: reference
lang: dart
description: SDK configuration example for Flutter.
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
```
