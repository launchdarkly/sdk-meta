---
id: flutter-client-sdk/sdk-docs/migration-3-to-4-connection-mode-changes-flutter-sdk-v4
sdk: flutter-client-sdk
kind: reference
lang: dart
description: "Flutter SDK v4 in section \"Connection mode changes\""
---

```dart
final config = LDConfig(
  CredentialSource.fromEnvironment(),
  AutoEnvAttributes.enabled,
  dataSourceConfig: DataSourceConfig(
    initialConnectionMode: ConnectionMode.offline // or .polling, or .streaming
  ),
);
```
