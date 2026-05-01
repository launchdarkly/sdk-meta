---
id: flutter-client-sdk/sdk-docs/migration-3-to-4-understanding-changes-to-ldconfig-flutter-sdk-v4
sdk: flutter-client-sdk
kind: reference
lang: dart
description: "Flutter SDK v4 in section \"Understanding changes to LDConfig\""
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
