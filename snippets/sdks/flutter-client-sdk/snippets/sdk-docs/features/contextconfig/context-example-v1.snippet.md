---
id: flutter-client-sdk/sdk-docs/features/contextconfig/context-example-v1
sdk: flutter-client-sdk
kind: reference
lang: dart
description: User example for Flutter SDK v1.x.
validation:
  scaffold: flutter-client-sdk/scaffolds/flutter-syntax-only-v2

---

```dart
LDUser user = LDUserBuilder('example-user-key')
    .email('sandy@example.com')
    .firstName('Sandy')
    .lastName('Smith')
    .custom('group', LDValue.ofString('microsoft'))
    .build();
```
