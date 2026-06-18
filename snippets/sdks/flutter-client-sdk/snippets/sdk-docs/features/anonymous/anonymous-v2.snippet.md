---
id: flutter-client-sdk/sdk-docs/features/anonymous/anonymous-v2
sdk: flutter-client-sdk
kind: reference
lang: dart
description: Anonymous context example for Flutter, SDK v2.x.
validation:
  scaffold: flutter-client-sdk/scaffolds/flutter-syntax-only-v2
---

```dart
LDContextBuilder builder = LDContextBuilder();
builder.kind('user', 'example-user-key')
    .anonymous(true);
LDContext context = builder.build();

```
