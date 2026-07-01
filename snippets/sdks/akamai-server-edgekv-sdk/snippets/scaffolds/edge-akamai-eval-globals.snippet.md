---
id: akamai-server-edgekv-sdk/scaffolds/edge-akamai-eval-globals
sdk: akamai-server-edgekv-sdk
kind: scaffold
lang: typescript
file: _globals.d.ts
description: |
  Ambient companion for `edge-akamai-eval`. Declares `ldClient` as the
  real return type of the Akamai SDK's `init`, so the evaluate fragment
  type-checks `ldClient.variation(...)` against the real client API
  without the scaffold having to inject a statement ahead of the
  fragment's own imports.
---

```typescript
declare const ldClient: ReturnType<typeof import('@launchdarkly/akamai-server-edgekv-sdk').init>;
```
