---
id: flutter-client-sdk/sdk-docs/features/contextconfig/multi-context-v4
sdk: flutter-client-sdk
kind: reference
lang: dart
description: Multi-context example for Flutter SDK v4.
validation:
  scaffold: flutter-client-sdk/scaffolds/flutter-syntax-only

---

```dart
LDContextBuilder builder = LDContextBuilder();
builder.kind('user', 'example-user-key')
  .name('Sandy');
builder.kind('organization', 'example-organization-key')
  .name('Global Health Services');
LDContext context = builder.build();
```
