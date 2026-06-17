---
id: roku-client-sdk/sdk-docs/features/events/track
sdk: roku-client-sdk
kind: reference
lang: brightscript
description: Custom event tracking example for Roku (BrightScript).
validation:
  scaffold: roku-client-sdk/scaffolds/roku-syntax-only

---

```brightscript
' without optional data
launchDarkly.track("example-event-key")

' with optional data
launchDarkly.track("example-event-key", {"customField": 123})

' with optional numeric metric
launchDarkly.track("example-event-key", invalid, 52.3)

```
