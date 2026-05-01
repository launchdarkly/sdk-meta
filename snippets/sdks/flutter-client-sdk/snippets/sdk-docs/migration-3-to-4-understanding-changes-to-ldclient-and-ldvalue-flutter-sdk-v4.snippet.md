---
id: flutter-client-sdk/sdk-docs/migration-3-to-4-understanding-changes-to-ldclient-and-ldvalue-flutter-sdk-v4
sdk: flutter-client-sdk
kind: reference
lang: dart
description: "Flutter SDK v4 in section \"Understanding changes to LDClient and LDValue\""
---

```dart
final config = LDConfig(
  CredentialSource.fromEnvironment(),
  AutoEnvAttributes.enabled,
  dataSourceConfig: DataSourceConfig(
    evaluationReasons: true
  ),
);

final context = LDContextBuilder()
  .kind("user", "example-user-key")
  .build();

final client = LDClient(config, context);
await client.start().timeout(const Duration(seconds: 30));
```
