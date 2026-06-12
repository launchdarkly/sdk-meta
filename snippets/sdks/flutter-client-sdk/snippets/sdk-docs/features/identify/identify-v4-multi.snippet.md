---
id: flutter-client-sdk/sdk-docs/features/identify/identify-v4-multi
sdk: flutter-client-sdk
kind: reference
lang: dart
description: Identify example for the Flutter SDK v4 (multi-context).
validation:
  scaffold: flutter-client-sdk/scaffolds/flutter-syntax-only
---

```dart
final updatedMultiContext = LDContextBuilder()
  .kind('user', 'example-user-key')
  .kind('device', 'example-device-key')
  .build();

await client.identify(updatedMultiContext);
```
