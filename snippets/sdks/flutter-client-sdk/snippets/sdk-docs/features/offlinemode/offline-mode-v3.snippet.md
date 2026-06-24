---
id: flutter-client-sdk/sdk-docs/features/offlinemode/offline-mode-v3
sdk: flutter-client-sdk
kind: reference
lang: dart
description: Offline mode example for Flutter SDK v3.x.
validation:
  scaffold: flutter-client-sdk/scaffolds/flutter-syntax-only-v3

---

```dart
LDConfig ldConfig = LDConfigBuilder('example-mobile-key', AutoEnvAttributes.Enabled)
    .offline(true)
    .build();

await LDClient.start(ldConfig, context);

// Or to switch an already-instantiated client to offline mode:
await LDClient.setOnline(false);
```
