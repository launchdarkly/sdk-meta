---
id: flutter-client-sdk/sdk-docs/features/privateattrs/context-v4
sdk: flutter-client-sdk
kind: reference
lang: dart
description: Marking context attributes private with the context builder for Flutter SDK v4.
validation:
  scaffold: flutter-client-sdk/scaffolds/flutter-syntax-only

---

```dart
final context = LDContextBuilder()
  .kind('user', 'example-user-key')
  .setString('name', 'Sandy')
  .setString('email', 'sandy@example.com', private: true)
  .setString('group', 'microsoft', private: true)
  .build();
```
