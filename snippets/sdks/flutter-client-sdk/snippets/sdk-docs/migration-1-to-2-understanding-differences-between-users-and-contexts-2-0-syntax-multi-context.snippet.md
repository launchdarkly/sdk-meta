---
id: flutter-client-sdk/sdk-docs/migration-1-to-2-understanding-differences-between-users-and-contexts-2-0-syntax-multi-context
sdk: flutter-client-sdk
kind: reference
lang: dart
description: "2.0 syntax, multi-context in section \"Understanding differences between users and contexts\""
---

```dart
LDContextBuilder builder = LDContextBuilder();
builder.kind('user', 'example-user-key');
builder.kind('device', 'example-device-key');
LDContext context = builder.build();
```
