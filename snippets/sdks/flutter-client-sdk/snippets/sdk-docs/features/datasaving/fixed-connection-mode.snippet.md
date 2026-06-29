---
id: flutter-client-sdk/sdk-docs/features/datasaving/fixed-connection-mode
sdk: flutter-client-sdk
kind: reference
lang: dart
description: Set a fixed connection mode with manual mode switching for Flutter.
validation:
  scaffold: flutter-client-sdk/scaffolds/flutter-syntax-only
---

```dart
final config = LDConfig(
  CredentialSource.fromEnvironment(),
  AutoEnvAttributes.enabled,
  dataSystem: const DataSystemConfig(
    initialConnectionMode: ConnectionModeId.polling,
  ),
);
```
