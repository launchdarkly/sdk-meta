---
id: flutter-client-sdk/sdk-docs/features/contextconfig/multi-context-v2
sdk: flutter-client-sdk
kind: reference
lang: dart
description: Multi-context example for Flutter SDK v2.x+.
validation:
  scaffold: flutter-client-sdk/scaffolds/flutter-syntax-only-v2

---

```dart
LDContextBuilder builder = LDContextBuilder();
builder.kind('user', 'example-user-key');
builder.kind('device', 'example-device-key');
LDContext context = builder.build();
```
