---
id: flutter-client-sdk/sdk-docs/features/logging/logging-v4
sdk: flutter-client-sdk
kind: reference
lang: dart
description: Logger configuration example for Flutter SDK v4.
validation:
  scaffold: flutter-client-sdk/scaffolds/flutter-syntax-only

---

```dart
final logger = LDLogger(level: LDLogLevel.warn);

final config = LDConfig(
  CredentialSource.fromEnvironment(),
  AutoEnvAttributes.enabled,
  logger: logger,
);
```
