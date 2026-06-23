---
id: flutter-client-sdk/sdk-docs/features/identify/identify-v3-multi
sdk: flutter-client-sdk
kind: reference
lang: dart
description: Identify example for the Flutter SDK v3.x (multi-context).
validation:
  scaffold: flutter-client-sdk/scaffolds/flutter-syntax-only-v3
---

```dart
LDContextBuilder builder = LDContextBuilder();
builder.kind('user', 'example-user-key');
builder.kind('device', 'example-device-key');

LDContext updatedMultiContext = builder.build();

await LDClient.identify(updatedMultiContext);
```
