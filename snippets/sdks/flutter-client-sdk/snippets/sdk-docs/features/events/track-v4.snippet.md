---
id: flutter-client-sdk/sdk-docs/features/events/track-v4
sdk: flutter-client-sdk
kind: reference
lang: dart
description: Custom event tracking example for Flutter SDK v4.
validation:
  scaffold: flutter-client-sdk/scaffolds/flutter-syntax-only

---

```dart
client.track('example-event-key', data: LDValue.buildObject().addBool("clicked-button", true).build());
```
