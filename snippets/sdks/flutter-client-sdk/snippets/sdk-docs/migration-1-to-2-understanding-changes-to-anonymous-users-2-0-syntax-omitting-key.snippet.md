---
id: flutter-client-sdk/sdk-docs/migration-1-to-2-understanding-changes-to-anonymous-users-2-0-syntax-omitting-key
sdk: flutter-client-sdk
kind: reference
lang: dart
description: "2.0 syntax, omitting key in section \"Understanding changes to anonymous users\""
---

```dart
LDContextBuilder builder = LDContextBuilder();
builder.kind('device')
    .set('os', LDValue.ofString('Android 25'))
    .set('device', LDValue.ofString('Pixel XL marlin'));
LDContext context = builder.build();
```
