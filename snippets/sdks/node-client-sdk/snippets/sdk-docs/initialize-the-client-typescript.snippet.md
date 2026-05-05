---
id: node-client-sdk/sdk-docs/initialize-the-client-typescript
sdk: node-client-sdk
kind: reference
lang: typescript
description: "TypeScript in section \"Initialize the client\""
validation:
  scaffold: node-client-sdk/scaffolds/node-client-syntax-only
---

```ts
client.on('initialized', () => {
  // initialization succeeded, flag values are now available
  const boolFlagValue = client.variation('example-flag-key', false) as boolean;
  const numberFlagValue = client.variation('example-flag-key', 2) as number;
  const stringFlagValue = client.variation('example-flag-key', 'default') as string;
  // etc.
});
```
