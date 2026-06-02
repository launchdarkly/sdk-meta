---
id: flutter-client-sdk/sdk-docs/initialize-the-client-flutter-sdk-v1-x-2
sdk: flutter-client-sdk
kind: reference
lang: dart
description: "Flutter SDK v1.x in section \"Initialize the client\""
# TODO(snippet-bug): body uses Flutter SDK v1.x API (LDConfigBuilder,
# LDUser, LDClient.start) — those types were removed in v2.x and
# replaced with LDConfig direct construction + LDContext. The
# flutter-syntax-only scaffold compiles against the latest
# launchdarkly_flutter_client_sdk, so this v1-shape call fails. Fix
# in the follow-up snippet-bugs PR: either update to current API or
# pin a v1 SDK in a parallel scaffold if back-compat docs must stay.
---

```dart
LDConfig config = LDConfigBuilder('example-mobile-key').build();
LDUser user = LDUserBuilder('example-user-key')
    .email('sandy@example.com')
    .build();

await LDClient.start(config, user);
await LDClient.startFuture(timeLimit: Duration(seconds: 5));
```
