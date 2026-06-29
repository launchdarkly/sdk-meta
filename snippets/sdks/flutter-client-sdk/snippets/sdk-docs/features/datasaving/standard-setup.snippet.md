---
id: flutter-client-sdk/sdk-docs/features/datasaving/standard-setup
sdk: flutter-client-sdk
kind: reference
lang: dart
description: Data saving mode standard setup for Flutter.
validation:
  scaffold: flutter-client-sdk/scaffolds/flutter-syntax-only
---

```dart
final config = LDConfig(
  CredentialSource.fromEnvironment(),
  AutoEnvAttributes.enabled,
  dataSystem: const DataSystemConfig(),
);
```
