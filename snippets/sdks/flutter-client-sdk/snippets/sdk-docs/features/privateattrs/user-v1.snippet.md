---
id: flutter-client-sdk/sdk-docs/features/privateattrs/user-v1
sdk: flutter-client-sdk
kind: reference
lang: dart
description: Marking user attributes private with the user builder for Flutter SDK v1.x.
validation:
  scaffold: flutter-client-sdk/scaffolds/flutter-syntax-only-v2

---

```dart
LDUser user = LDUserBuilder('example-user-key')
    .name('Sandy')
    .privateEmail('sandy@example.com')
    .privateCustom('group', LDValue.ofString('microsoft'))
    .build();
```
