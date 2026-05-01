---
id: flutter-client-sdk/sdk-docs/migration-1-to-2-working-with-built-in-and-custom-attributes-2-0-syntax-context-with-attributes
sdk: flutter-client-sdk
kind: reference
lang: dart
description: "2.0 syntax, context with attributes in section \"Working with built-in and custom attributes\""
---

```dart
LDContextBuilder builder = LDContextBuilder();
builder.kind('user', 'example-user-key')
    .set('email', LDValue.ofString('sandy@example.com'))
    .set('name', LDValue.ofString('Sandy'))
    .set('age', LDValue.ofNum(32))
    .set('group', LDValue.ofString('Global Health Services'));
LDContext context = builder.build();
```
