---
id: flutter-client-sdk/sdk-docs/features/anonymous/anonymous-v1x
sdk: flutter-client-sdk
kind: reference
lang: dart
description: Anonymous user example for Flutter, SDK v1.x.
validation:
  scaffold: flutter-client-sdk/scaffolds/flutter-syntax-only-v2
---

```dart
LDUser user = LDUserBuilder('example-user-key')
    .anonymous(true)
    .build();
```
