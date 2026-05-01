---
id: flutter-client-sdk/sdk-docs/migration-3-to-4-understanding-changes-to-ldclient-and-ldvalue-flutter-sdk-v3-x
sdk: flutter-client-sdk
kind: reference
lang: dart
description: "Flutter SDK v3.x in section \"Understanding changes to LDClient and LDValue\""
---

```dart
LDConfig ldConfig = LDConfigBuilder('example-mobile-key', AutoEnvAttributes.Enabled)
    .evaluationReasons(true)
    .build();

LDContextBuilder builder = LDContextBuilder();
builder.kind('user', 'example-user-key');

LDContext context = builder.build();
await LDClient.start(config, context);
```
