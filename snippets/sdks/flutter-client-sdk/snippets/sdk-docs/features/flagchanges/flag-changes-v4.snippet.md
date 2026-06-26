---
id: flutter-client-sdk/sdk-docs/features/flagchanges/flag-changes-v4
sdk: flutter-client-sdk
kind: reference
lang: dart
description: Flag change subscription example for Flutter SDK v4.
validation:
  scaffold: flutter-client-sdk/scaffolds/flutter-syntax-only

---

```dart
final sub = client.flagChanges.listen((changeEvent) {
  for(var flagKey in changeEvent.keys) {
    print(client.jsonVariation(flagKey, LDValue.ofString('default')));
  }
});
```
