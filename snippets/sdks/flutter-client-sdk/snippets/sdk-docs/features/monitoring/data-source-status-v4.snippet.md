---
id: flutter-client-sdk/sdk-docs/features/monitoring/data-source-status-v4
sdk: flutter-client-sdk
kind: reference
lang: dart
description: Data source status access and change listener for Flutter SDK v4.
validation:
  scaffold: flutter-client-sdk/scaffolds/flutter-syntax-only

---

```dart
// get the current status
final status = client.dataSourceStatus;

// listen for changes
final sub = client.dataSourceStatusChanges.listen((status){
  // act on status
});
```
