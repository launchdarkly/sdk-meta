---
id: flutter-client-sdk/sdk-docs/migration-3-to-4-logging-flutter-sdk-v4
sdk: flutter-client-sdk
kind: reference
lang: dart
description: "Flutter SDK v4 in section \"Logging\""
---

```dart
final logger = LDLogger(level: LDLogLevel.warn);

final config = LDConfig(
  CredentialSource.fromEnvironment(),
  AutoEnvAttributes.enabled,
  logger: logger,
);
```
