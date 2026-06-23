---
id: roku-client-sdk/sdk-docs/features/contextconfig/multi-context-v2
sdk: roku-client-sdk
kind: reference
lang: brightscript
description: Multi-context example for Roku SDK v2.0.
validation:
  scaffold: roku-client-sdk/scaffolds/roku-syntax-only

---

```brightscript
context = LaunchDarklyCreateContext({
    "kind": "multi",
    "user": { "key": "example-user-key", "name": "Sandy" },
    "org": { "key": "org-key-789xyz", "name": "LaunchDarkly" }
})
```
