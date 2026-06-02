---
id: haskell-server-sdk/sdk-docs/evaluate-a-context-haskell-sdk-v3-x
sdk: haskell-server-sdk
kind: reference
lang: haskell
description: "Haskell SDK v3.x in section \"Evaluate a context\""
# TODO(snippet-bug): body uses Haskell SDK v3.x API (`makeUser`),
# replaced by `makeContext` in v4.0. The haskell-syntax-only scaffold
# compiles against the latest launchdarkly-server-sdk, so `makeUser`
# fails name resolution. Fix in the follow-up snippet-bugs PR: either
# update to current `makeContext` API and drop the v3.x-specific
# snippet, or pin a v3.x SDK in a parallel scaffold if back-compat
# docs must stay live.
---

```haskell
boolVariation client "example-flag-key" (makeUser "example-user-key") False
```
