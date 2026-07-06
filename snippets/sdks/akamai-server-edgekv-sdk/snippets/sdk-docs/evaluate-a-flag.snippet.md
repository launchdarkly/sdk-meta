---
id: akamai-server-edgekv-sdk/sdk-docs/evaluate-a-flag
sdk: akamai-server-edgekv-sdk
kind: reference
lang: typescript
file: akamai-server-edgekv-sdk/sdk-docs/evaluate-a-flag.ts
description: "Akamai edge SDK in section \"Evaluate a flag\""
validation:
  scaffold: akamai-server-edgekv-sdk/scaffolds/edge-akamai-eval
---

```ts
import { LDContext } from '@launchdarkly/akamai-server-edgekv-sdk';

const context: LDContext = {
  kind: 'org',
  key: 'example-organization-key',
  someAttribute: 'example-attribute-value',
};

const flagValue = await ldClient.variation('example-flag-key', context, false);
```
