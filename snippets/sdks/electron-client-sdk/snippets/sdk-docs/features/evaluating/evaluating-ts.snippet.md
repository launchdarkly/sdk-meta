---
id: electron-client-sdk/sdk-docs/features/evaluating/evaluating-ts
sdk: electron-client-sdk
kind: reference
lang: typescript
description: Flag evaluation example for Electron.
validation:
  scaffold: electron-client-sdk/scaffolds/electron-syntax-only

---

```typescript
const boolFlagValue = client.boolVariation('example-bool-flag-key', false);
const numberFlagValue = client.numberVariation('example-number-flag-key', 2);
const stringFlagValue = client.stringVariation('example-string-flag-key', 'default');

// proceed based on flag value, for example:

if (boolFlagValue)  {
  // feature flag targeting is on
} else {
  // feature flag targeting is off
}
```
