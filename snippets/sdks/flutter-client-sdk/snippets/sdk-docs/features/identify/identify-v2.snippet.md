---
id: flutter-client-sdk/sdk-docs/features/identify/identify-v2
sdk: flutter-client-sdk
kind: reference
lang: dart
description: Identify example for the Flutter SDK v2.x (single context).
validation:
  scaffold: flutter-client-sdk/scaffolds/flutter-syntax-only-v2
---

```dart
LDContextBuilder builder = LDContextBuilder();
builder.kind('user', 'example-user-key')
    .set('email', LDValue.ofString('sandy@example.com'));
LDContext updatedContext = builder.build();

await LDClient.identifyWithContext(updatedContext);
```
