---
id: js-client-sdk/sdk-docs/features/evaluation-reasons/evaluation-reasons-v4
sdk: js-client-sdk
kind: reference
lang: typescript
description: Flag evaluation reason example for JavaScript SDK v4.x.
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only

---

```typescript
const options = { evaluationReasons: true };
const client = createClient('example-client-side-id', context, options);
client.start();

await client.waitForInitialization({ timeout: 5 });

const detail = client.boolVariationDetail('example-flag-key', false);
// or stringVariationDetail for a string-valued flag, and so on.

const value = detail.value;
const index = detail.variationIndex;
const reason = detail.reason;
```
