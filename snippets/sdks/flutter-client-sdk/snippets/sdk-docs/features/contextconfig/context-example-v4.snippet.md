---
id: flutter-client-sdk/sdk-docs/features/contextconfig/context-example-v4
sdk: flutter-client-sdk
kind: reference
lang: dart
description: Context example for Flutter SDK v4.
validation:
  scaffold: flutter-client-sdk/scaffolds/flutter-syntax-only

---

```dart
final context = LDContextBuilder()
  .kind('user', 'example-user-key')
  .setString('email', 'sandy@example.com')
  .setString('firstName', 'Sandy')
  .setString('lastName', 'Smith')
  .setString('group', 'microsoft')
  .build();
```
