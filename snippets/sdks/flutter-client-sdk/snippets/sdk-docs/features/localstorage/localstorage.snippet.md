---
id: flutter-client-sdk/sdk-docs/features/localstorage/localstorage
sdk: flutter-client-sdk
kind: reference
lang: dart
description: Local storage caching example for Flutter.
validation:
  scaffold: flutter-client-sdk/scaffolds/flutter-syntax-only
---

```dart
// Local storage is enabled by default
// You can optionally configure the maximum number of cached contexts (default is 5)
final config = LDConfig(
  'example-mobile-key',
  AutoEnvAttributes.enabled,
  persistence: PersistenceConfig(maxCachedContexts: 3),
);

final context = LDContextBuilder()
  .kind('user', 'example-context-key')
  .build();

final client = LDClient(config, context);
await client.start();
```
