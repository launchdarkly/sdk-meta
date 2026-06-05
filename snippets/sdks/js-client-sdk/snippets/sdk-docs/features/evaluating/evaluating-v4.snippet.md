---
id: js-client-sdk/sdk-docs/features/evaluating/evaluating-v4
sdk: js-client-sdk
kind: reference
lang: typescript
description: Typed flag evaluation example for JavaScript SDK v4.x.
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only

---

```typescript
const boolFlagValue = client.boolVariation('example-bool-flag-key', false);
const numberFlagValue = client.numberVariation('example-numeric-flag-key', 2);
const stringValue = client.stringVariation('example-string-flag-key', 'default');
const jsonValue = client.jsonVariation('example-json-flag-key', '{}');
```
