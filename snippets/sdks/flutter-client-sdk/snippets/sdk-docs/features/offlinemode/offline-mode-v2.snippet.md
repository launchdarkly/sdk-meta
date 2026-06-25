---
id: flutter-client-sdk/sdk-docs/features/offlinemode/offline-mode-v2
sdk: flutter-client-sdk
kind: reference
lang: dart
description: Offline mode example for Flutter SDK v2.x.
validation:
  scaffold: flutter-client-sdk/scaffolds/flutter-syntax-only-v2

---

```dart
LDConfig ldConfig = LDConfigBuilder('example-mobile-key')
    .offline(true)
    .build();

await LDClient.startWithContext(ldConfig, context);

// Or to switch an already-instantiated client to offline mode:
await LDClient.setOnline(false);
```
