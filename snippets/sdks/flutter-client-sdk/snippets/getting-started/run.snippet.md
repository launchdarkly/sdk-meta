---
id: flutter-client-sdk/getting-started/run
sdk: flutter-client-sdk
kind: run
lang: bash
description: Run the Flutter app with the mobile key passed via --dart-define.
inputs:
  mobileKey:
    type: mobile-key
    description: Mobile key embedded in the rendered Run command.
ld-application:
  slot: run
---

Run:

```bash
flutter run --dart-define LAUNCHDARKLY_MOBILE_KEY={{ mobileKey }}
```
