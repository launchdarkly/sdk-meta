---
id: roku-client-sdk/sdk-docs/features/identify/identify-v2
sdk: roku-client-sdk
kind: reference
lang: brightscript
description: Identify example for the Roku SDK v2.0 (BrightScript).
validation:
  scaffold: roku-client-sdk/scaffolds/roku-syntax-only
---

```brightscript
' before the end user logs in
device = {
  "kind": "device",
  "key": "example-device-key",
  "type": "tablet"
}
context = LaunchDarklyCreateContext(device)

' after the end user logs in
multi = {
  "kind": "multi",
  "device": device,
  "user": {"key": "example-user-key", "name": "Sandy"},
  "organization": {"key": "example-organization-key", "name": "Acme, Inc."}
}

launchDarkly.identify(multi)

```
