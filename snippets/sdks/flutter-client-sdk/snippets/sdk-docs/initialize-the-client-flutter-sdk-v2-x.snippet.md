---
id: flutter-client-sdk/sdk-docs/initialize-the-client-flutter-sdk-v2-x
sdk: flutter-client-sdk
kind: reference
lang: dart
description: "Flutter SDK v2.x in section \"Initialize the client\""
validation:
  scaffold: flutter-client-sdk/scaffolds/flutter-syntax-only-v2
---

```dart
LDConfig config = LDConfigBuilder('example-mobile-key').build();

// You'll need this context later, but you can ignore it for now.
LDContextBuilder builder = LDContextBuilder();
builder.kind('user', 'example-user-key');

LDContext context = builder.build();
await LDClient.startWithContext(config, context);
```
