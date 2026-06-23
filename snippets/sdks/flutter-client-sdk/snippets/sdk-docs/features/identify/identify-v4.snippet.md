---
id: flutter-client-sdk/sdk-docs/features/identify/identify-v4
sdk: flutter-client-sdk
kind: reference
lang: dart
description: Identify example for the Flutter SDK v4 (single context).
validation:
  scaffold: flutter-client-sdk/scaffolds/flutter-syntax-only
---

```dart
final updatedContext = LDContextBuilder()
  .kind('user', 'example-user-key')
  .setString('email', 'sandy@example.com')
  .build();

await client.identify(updatedContext);
```
