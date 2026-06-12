---
id: flutter-client-sdk/sdk-docs/features/privateattrs/context-v2
sdk: flutter-client-sdk
kind: reference
lang: dart
description: Marking context attributes private with the context builder for Flutter SDK v2.x+.
validation:
  scaffold: flutter-client-sdk/scaffolds/flutter-syntax-only-v2

---

```dart
LDContextBuilder builder = LDContextBuilder();
builder.kind('user', 'example-user-key')
    .name('Sandy')
    .set('email', LDValue.ofString('sandy@example.com'))
    .set('group', LDValue.ofString('microsoft'))
    .privateAttributes(['email', 'group']);
LDContext context = builder.build();
```
