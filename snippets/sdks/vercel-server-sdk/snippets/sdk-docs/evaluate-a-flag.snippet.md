---
id: vercel-server-sdk/sdk-docs/evaluate-a-flag
sdk: vercel-server-sdk
kind: reference
lang: typescript
file: vercel-server-sdk/sdk-docs/evaluate-a-flag.ts
description: "Vercel edge SDK in section \"Evaluate a flag\""
validation:
  scaffold: vercel-server-sdk/scaffolds/edge-vercel-eval
---

```typescript
const ldContext = {
  kind: 'org',
  key: 'example-organization-key',
  someAttribute: 'example-attribute-value',
}
const flagValue = await ldClient.variation('example-flag-key', ldContext, true)
```
