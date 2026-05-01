---
id: roku-client-sdk/sdk-docs/migration-1-to-2-understanding-changes-to-alias-events-2-0-syntax-associating-two-contexts
sdk: roku-client-sdk
kind: reference
lang: brightscript
description: "2.0 syntax, associating two contexts in section \"Understanding changes to alias events\""
---

```brightscript
attributes = {
    "kind": "multi",
    "user": {
        "key": "example-user-key"
    },
    "device": {
        "key": "example-device-key"
    }
}
multiContext = LaunchDarklyCreateContext(attributes)
client.identify(multiContext)
```
