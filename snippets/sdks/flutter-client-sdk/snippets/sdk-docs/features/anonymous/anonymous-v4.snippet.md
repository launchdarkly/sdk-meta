---
id: flutter-client-sdk/sdk-docs/features/anonymous/anonymous-v4
sdk: flutter-client-sdk
kind: reference
lang: dart
description: Anonymous context example for Flutter, SDK v4.
validation:
  scaffold: flutter-client-sdk/scaffolds/flutter-syntax-only
---

```dart
final context = LDContextBuilder()
  .kind('user', 'example-user-key')
  .anonymous(true)
  .build();
```
