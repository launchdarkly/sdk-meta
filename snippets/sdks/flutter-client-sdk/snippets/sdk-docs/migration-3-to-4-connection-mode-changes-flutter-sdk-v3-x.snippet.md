---
id: flutter-client-sdk/sdk-docs/migration-3-to-4-connection-mode-changes-flutter-sdk-v3-x
sdk: flutter-client-sdk
kind: reference
lang: dart
description: "Flutter SDK v3.x in section \"Connection mode changes\""
---

```dart
LDConfig ldConfig = LDConfigBuilder('example-mobile-key', AutoEnvAttributes.Enabled)
    .offline(true)
    .build();

await LDClient.start(ldConfig, context);
```
