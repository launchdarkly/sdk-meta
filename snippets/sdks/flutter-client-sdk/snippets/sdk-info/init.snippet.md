---
id: flutter-client-sdk/sdk-info/init
sdk: flutter-client-sdk
kind: init
lang: dart
file: flutter-client-sdk/init.txt
description: |
  Client initialization snippet for flutter-client-sdk (v4 API).
  Compile-checked against the real SDK via the flutter-syntax-only
  scaffold; the harness lifts the import to module scope.
validation:
  scaffold: flutter-client-sdk/scaffolds/flutter-syntax-only
---

```dart
import 'package:launchdarkly_flutter_client_sdk/launchdarkly_flutter_client_sdk.dart';

// This is your LaunchDarkly mobile key.
// Never hardcode your mobile key in production.
final config = LDConfig('YOUR_MOBILE_KEY', AutoEnvAttributes.enabled);
final context = LDContextBuilder().kind('user', 'example-user-key').build();

final client = LDClient(config, context);

// Wait up to five seconds for the client to connect to LaunchDarkly.
await client.start().timeout(const Duration(seconds: 5));
```
