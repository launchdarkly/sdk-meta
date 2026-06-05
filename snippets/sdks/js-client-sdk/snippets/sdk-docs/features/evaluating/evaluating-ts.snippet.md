---
id: js-client-sdk/sdk-docs/features/evaluating/evaluating-ts
sdk: js-client-sdk
kind: reference
lang: typescript
description: Flag evaluation example for JavaScript.
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only

---

```typescript
const boolFlagValue = client.variation('example-bool-flag-key', false) as boolean;
const numberFlagValue = client.variation('example-numeric-flag-key', 2) as number;
const stringFlagValue = client.variation('example-string-flag-key', 'default') as string;
```
