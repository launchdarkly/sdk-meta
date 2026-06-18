---
id: flutter-client-sdk/sdk-docs/features/identify/identify-v1x
sdk: flutter-client-sdk
kind: reference
lang: dart
description: Identify example for the Flutter SDK v1.x.
validation:
  scaffold: flutter-client-sdk/scaffolds/flutter-syntax-only-v2
---

```dart
LDUser updatedUser = LDUserBuilder('user key')
    .email('sandy@example.com')
    .build();

await LDClient.identify(updatedUser);
```
