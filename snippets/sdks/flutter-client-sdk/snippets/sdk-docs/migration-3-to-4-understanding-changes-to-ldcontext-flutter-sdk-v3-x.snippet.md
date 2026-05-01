---
id: flutter-client-sdk/sdk-docs/migration-3-to-4-understanding-changes-to-ldcontext-flutter-sdk-v3-x
sdk: flutter-client-sdk
kind: reference
lang: dart
description: "Flutter SDK v3.x in section \"Understanding changes to LDContext\""
---

```dart
LDContextBuilder builder = LDContextBuilder();
builder.kind('user', 'example-user-key')
  .set('name', LDValue.ofString('Sandy Smith'))
  .set('employeeID', LDValue.ofNum(1234))
  .set('fullTimeEmployee', LDValue.ofBool(true))
  .privateAttributes(['employeeID', 'fullTimeEmployee']);
LDContext context = builder.build();
```
