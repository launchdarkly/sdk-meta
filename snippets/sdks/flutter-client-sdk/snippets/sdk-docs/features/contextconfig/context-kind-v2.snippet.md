---
id: flutter-client-sdk/sdk-docs/features/contextconfig/context-kind-v2
sdk: flutter-client-sdk
kind: reference
lang: dart
description: Context with a non-user kind for Flutter SDK v2.x+.
validation:
  scaffold: flutter-client-sdk/scaffolds/flutter-syntax-only-v2

---

```dart
LDContextBuilder builder = LDContextBuilder();
builder.kind('device', 'example-device-key');
LDContext context = builder.build();
```
