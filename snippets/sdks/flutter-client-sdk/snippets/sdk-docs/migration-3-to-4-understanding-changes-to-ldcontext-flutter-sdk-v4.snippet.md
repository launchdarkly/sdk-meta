---
id: flutter-client-sdk/sdk-docs/migration-3-to-4-understanding-changes-to-ldcontext-flutter-sdk-v4
sdk: flutter-client-sdk
kind: reference
lang: dart
description: "Flutter SDK v4 in section \"Understanding changes to LDContext\""
---

```dart
final context = LDContextBuilder()
  .kind('user', 'example-user-key')
  .setString('name', 'Sandy Smith') // or .name('Sandy Smith')
  .setNum('employeeID', 1234, private: true)
  .setBool('fullTimeEmployee', true, private: true)
  .build();
```
