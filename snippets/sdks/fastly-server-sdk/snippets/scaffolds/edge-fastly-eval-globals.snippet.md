---
id: fastly-server-sdk/scaffolds/edge-fastly-eval-globals
sdk: fastly-server-sdk
kind: scaffold
lang: typescript
file: _globals.d.ts
description: |
  Ambient companion for edge-fastly-eval. Declares `ldClient` as the real
  return type of the Fastly SDK's init so the evaluate fragment
  type-checks variation(...) against the real client API.
---

```typescript
declare const ldClient: ReturnType<typeof import('@launchdarkly/fastly-server-sdk').init>;
```
