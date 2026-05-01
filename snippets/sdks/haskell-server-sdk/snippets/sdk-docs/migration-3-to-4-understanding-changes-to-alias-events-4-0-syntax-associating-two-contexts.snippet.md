---
id: haskell-server-sdk/sdk-docs/migration-3-to-4-understanding-changes-to-alias-events-4-0-syntax-associating-two-contexts
sdk: haskell-server-sdk
kind: reference
lang: haskell
description: "4.0 syntax, associating two contexts in section \"Understanding changes to alias events\""
---

```haskell
let context1 = makeContext "example-user-key" "user"
    context2 = makeContext "example-device-key" "device"
    multiContext = makeMultiContext [context1, context2]
 in identify client multiContext
```
