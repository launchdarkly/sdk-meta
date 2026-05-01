---
id: flutter-client-sdk/sdk-docs/migration-3-to-4-subscribing-to-flag-changes-flutter-sdk-v4
sdk: flutter-client-sdk
kind: reference
lang: dart
description: "Flutter SDK v4 in section \"Subscribing to flag changes\""
---

```dart
final sub = client.flagChanges.listen((changeEvent) {
  for(var flagKey in changeEvent.keys) {
    print(client.jsonVariation(flagKey, LDValue.ofString('default')));
  }
});
```
