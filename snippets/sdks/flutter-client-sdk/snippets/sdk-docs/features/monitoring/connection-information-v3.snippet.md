---
id: flutter-client-sdk/sdk-docs/features/monitoring/connection-information-v3
sdk: flutter-client-sdk
kind: reference
lang: dart
description: Connection information retrieval for Flutter SDK v3.x.
validation:
  scaffold: flutter-client-sdk/scaffolds/flutter-syntax-only-v3

---

```dart
LDConnectionInformation? connectionInfo = await LDClient.getConnectionInformation();
// The current connection state
LDConnectionState? connectionState = connectionInfo?.connectionState;
// Most recent successful flag cache update
DateTime? lastSuccess = connectionInfo?.lastSuccessfulConnection;
// Most recent unsuccessful flag cache update attempt
DateTime? lastfailure = connectionInfo?.lastFailedConnection;
// Most recent failure or null
LDFailure? ldFailure = connectionInfo?.lastFailure;
```
