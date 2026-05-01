---
id: flutter-client-sdk/sdk-docs/migration-1-to-2-working-with-built-in-and-custom-attributes-1-x-syntax-user-with-attributes
sdk: flutter-client-sdk
kind: reference
lang: dart
description: "1.x syntax, user with attributes in section \"Working with built-in and custom attributes\""
---

```dart
LDUser user = LDUserBuilder('example-user-key')
    .email('sandy@example.com')
    .firstName('Sandy')
    .lastName('Smith')
    .custom('group', LDValue.ofString('Global Health Services'))
    .build();
```
