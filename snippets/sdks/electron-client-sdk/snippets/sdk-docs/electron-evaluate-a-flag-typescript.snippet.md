---
id: electron-client-sdk/sdk-docs/electron-evaluate-a-flag-typescript
sdk: electron-client-sdk
kind: reference
lang: typescript
description: "TypeScript in section \"Evaluate a flag\""
---

```ts
const boolFlagValue = client.variation('bool-example-flag-key', false) as boolean;
const numberFlagValue = client.variation('number-example-flag-key', 2) as number;
const stringFlagValue = client.variation('string-example-flag-key', 'default') as string;

// proceed based on flag value, for example:

if (boolFlagValue)  {
  // feature flag targeting is on
} else {
  // feature flag targeting is off
}
```
