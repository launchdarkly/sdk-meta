---
id: roku-client-sdk/sdk-docs/features/privateattrs/context
sdk: roku-client-sdk
kind: reference
lang: brightscript
description: Marking context attributes private for Roku (BrightScript).
validation:
  scaffold: roku-client-sdk/scaffolds/roku-syntax-only

---

```brightscript
' when creating a context
context = LaunchDarklyCreateContext({
    "kind": "user",
    "key": "context-key-123-abc",
    "email": "sandy@example.com",
    "_meta": { privateAttributes: ["email"] }
})
```
