---
id: flutter-client-sdk/sdk-docs/features/events/track-v3
sdk: flutter-client-sdk
kind: reference
lang: dart
description: Custom event tracking example for Flutter SDK v3.x.
validation:
  scaffold: flutter-client-sdk/scaffolds/flutter-syntax-only-v3

---

```dart
await client.track('example-event-key', data: LDValue.buildObject().addBool("clicked-button", true).build());
```
