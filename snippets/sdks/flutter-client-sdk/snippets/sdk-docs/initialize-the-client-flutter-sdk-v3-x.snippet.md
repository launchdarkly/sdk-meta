---
id: flutter-client-sdk/sdk-docs/initialize-the-client-flutter-sdk-v3-x
sdk: flutter-client-sdk
kind: reference
lang: dart
description: "Flutter SDK v3.x in section \"Initialize the client\""
---

```dart
LDConfig config = LDConfigBuilder('example-mobile-key', AutoEnvAttributes.Enabled)
  .build();

// You'll need this context later, but you can ignore it for now.
LDContextBuilder builder = LDContextBuilder();
builder.kind('user', 'example-user-key');

LDContext context = builder.build();
await LDClient.start(config, context);
```
