---
id: flutter-client-sdk/sdk-docs/initialize-the-client-flutter-sdk-v4
sdk: flutter-client-sdk
kind: reference
lang: dart
description: "Flutter SDK v4 in section \"Initialize the client\""
---

```dart
final config = LDConfig(
  CredentialSource.fromEnvironment(),
  AutoEnvAttributes.enabled,
  // all other configuration options are optional
);

// You'll need this context later, but you can ignore it for now.
final context = LDContextBuilder()
  .kind("user", "example-user-key")
  .build();

final client = LDClient(config, context);
await client.start();
```
