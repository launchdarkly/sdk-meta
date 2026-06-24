---
id: flutter-client-sdk/sdk-docs/features/offlinemode/offline-mode-v1x
sdk: flutter-client-sdk
kind: reference
lang: dart
description: Offline mode example for Flutter SDK v1.x.
# The v1.x API surface used here (LDConfigBuilder with a single
# positional key, LDClient.start(config, user), LDClient.setOnline)
# still exists in the 2.x SDK (start is deprecated there in favor of
# startWithContext), so this fragment compiles against the
# flutter-client-v2 validator's pinned 2.x package.
validation:
  scaffold: flutter-client-sdk/scaffolds/flutter-syntax-only-v2

---

```dart
LDConfig ldConfig = LDConfigBuilder('example-mobile-key')
    .offline(true)
    .build();

await LDClient.start(ldConfig, user);

// Or to switch an already-instantiated client to offline mode:
await LDClient.setOnline(false);
```
