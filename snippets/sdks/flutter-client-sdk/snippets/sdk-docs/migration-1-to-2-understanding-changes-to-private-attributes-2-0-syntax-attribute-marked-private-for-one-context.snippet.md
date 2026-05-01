---
id: flutter-client-sdk/sdk-docs/migration-1-to-2-understanding-changes-to-private-attributes-2-0-syntax-attribute-marked-private-for-one-context
sdk: flutter-client-sdk
kind: reference
lang: dart
description: "2.0 syntax, attribute marked private for one context in section \"Understanding changes to private attributes\""
---

```dart
LDContextBuilder builder = LDContextBuilder();
builder.kind('user', 'example-user-key')
    .name('Sandy')
    .set('email', LDValue.ofString('sandy@example.com'))
    .set('group', LDValue.ofString('microsoft'))
    .privateAttributes(['name', 'group']);
LDContext context = builder.build();
```
