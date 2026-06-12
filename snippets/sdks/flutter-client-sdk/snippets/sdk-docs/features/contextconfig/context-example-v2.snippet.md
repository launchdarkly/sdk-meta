---
id: flutter-client-sdk/sdk-docs/features/contextconfig/context-example-v2
sdk: flutter-client-sdk
kind: reference
lang: dart
description: Context example for Flutter SDK v2.x+.
validation:
  scaffold: flutter-client-sdk/scaffolds/flutter-syntax-only-v2

---

```dart
LDContextBuilder builder = LDContextBuilder();
builder.kind('user', 'example-user-key')
    .set('email', LDValue.ofString('sandy@example.com'))
    .set('firstName', LDValue.ofString('Sandy'))
    .set('lastName', LDValue.ofString('Smith'))
    .set('group', LDValue.ofString('microsoft'));

LDContext context = builder.build();
```
