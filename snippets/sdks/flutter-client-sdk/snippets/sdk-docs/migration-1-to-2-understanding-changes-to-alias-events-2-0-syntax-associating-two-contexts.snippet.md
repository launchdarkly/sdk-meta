---
id: flutter-client-sdk/sdk-docs/migration-1-to-2-understanding-changes-to-alias-events-2-0-syntax-associating-two-contexts
sdk: flutter-client-sdk
kind: reference
lang: dart
description: "2.0 syntax, associating two contexts in section \"Understanding changes to alias events\""
---

```dart
LDContextBuilder builder = LDContextBuilder();
builder.kind('user', 'example-user-key');
builder.kind('device', 'example-device-key');

LDContext updatedMultiContext = builder.build();

await LDClient.identifyWithContext(updatedMultiContext);
```
