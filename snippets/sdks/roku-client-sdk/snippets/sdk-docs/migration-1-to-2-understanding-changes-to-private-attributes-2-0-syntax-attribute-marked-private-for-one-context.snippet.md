---
id: roku-client-sdk/sdk-docs/migration-1-to-2-understanding-changes-to-private-attributes-2-0-syntax-attribute-marked-private-for-one-context
sdk: roku-client-sdk
kind: reference
lang: brightscript
description: "2.0 syntax, attribute marked private for one context in section \"Understanding changes to private attributes\""
---

```brightscript
context = LaunchDarklyCreateContext({
    "kind": "user",
    "key": "context-key-123-abc",
    "email": "sandy@example.com",
    "_meta": { privateAttributes: ["email"] }
})
```
