---
id: flutter-client-sdk/sdk-docs/initialize-the-client-flutter-sdk-v1-x-2
sdk: flutter-client-sdk
kind: reference
lang: dart
description: "Flutter SDK v1.x in section \"Initialize the client\""
validation:
  scaffold: flutter-client-sdk/scaffolds/flutter-syntax-only
---

```dart
LDConfig config = LDConfigBuilder('example-mobile-key').build();
LDUser user = LDUserBuilder('example-user-key')
    .email('sandy@example.com')
    .build();

await LDClient.start(config, user);
await LDClient.startFuture(timeLimit: Duration(seconds: 5));
```
