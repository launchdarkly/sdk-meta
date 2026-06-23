---
id: flutter-client-sdk/sdk-docs/features/privateattrs/config-v4
sdk: flutter-client-sdk
kind: reference
lang: dart
description: Private attribute configuration for Flutter SDK v4.
validation:
  scaffold: flutter-client-sdk/scaffolds/flutter-syntax-only

---

```dart
final config = LDConfig(
  CredentialSource.fromEnvironment(),
  AutoEnvAttributes.enabled,
  allAttributesPrivate: true, // all attributes marked private
  globalPrivateAttributes: ['user/email', 'user/group'], // two attributes marked private for the 'user' context kind
);
```
