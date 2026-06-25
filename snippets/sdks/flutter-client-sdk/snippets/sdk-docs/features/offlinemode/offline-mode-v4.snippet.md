---
id: flutter-client-sdk/sdk-docs/features/offlinemode/offline-mode-v4
sdk: flutter-client-sdk
kind: reference
lang: dart
description: Offline mode example for Flutter SDK v4.
validation:
  scaffold: flutter-client-sdk/scaffolds/flutter-syntax-only

---

```dart
final config = LDConfig(
  CredentialSource.fromEnvironment(),
  AutoEnvAttributes.enabled,
  dataSourceConfig: DataSourceConfig(
    initialConnectionMode: ConnectionMode.streaming // or .polling, or .offline
  ),
);

final client = LDClient(config, context);
await client.start();

// To switch an already-instantiated client to offline mode:
client.offline = true;

// To switch it back:
client.offline = false;
```
