---
id: flutter-client-sdk/sdk-docs/initialize-the-client-flutter-sdk-v1-x
sdk: flutter-client-sdk
kind: reference
lang: dart
description: "Flutter SDK v1.x in section \"Initialize the client\""
---

```dart
LDConfig config = LDConfigBuilder('example-mobile-key').build();

// You'll need this context later, but you can ignore it for now.
LDUser user = LDUserBuilder('example-user-key')
    .email('sandy@example.com')
    .build();

await LDClient.start(config, user);
```
