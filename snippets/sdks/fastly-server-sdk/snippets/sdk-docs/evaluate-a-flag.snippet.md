---
id: fastly-server-sdk/sdk-docs/evaluate-a-flag
sdk: fastly-server-sdk
kind: reference
lang: typescript
file: fastly-server-sdk/sdk-docs/evaluate-a-flag.ts
description: "Fastly edge SDK in section \"Evaluate a flag\""
validation:
  scaffold: fastly-server-sdk/scaffolds/edge-fastly-eval
---

```ts
import type { LDContext } from '@launchdarkly/js-server-sdk-common';

const ldContext: LDContext = {
  kind: 'org',
  key: 'example-organization-key',
  someAttribute: 'example-attribute-value',
}
const flagValue = await ldClient.variation('example-flag-key', ldContext, false)
```
