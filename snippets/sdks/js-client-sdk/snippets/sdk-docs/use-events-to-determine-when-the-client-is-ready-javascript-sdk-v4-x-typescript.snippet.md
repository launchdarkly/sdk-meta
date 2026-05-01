---
id: js-client-sdk/sdk-docs/use-events-to-determine-when-the-client-is-ready-javascript-sdk-v4-x-typescript
sdk: js-client-sdk
kind: reference
lang: typescript
description: "JavaScript SDK, v4.x (TypeScript) in section \"Use events to determine when the client is ready\""
---

```ts
client.on('ready', () => {
  // initialization succeeded, flag values are now available
  //
  // in v4.x of the JavaScript SDK you can also use typed methods,
  // for example, `boolVariationDetail` for boolean feature flags
  const boolFlagValue = client.boolVariation('example-flag-key', false);
  const numberFlagValue = client.numberVariation('example-flag-key', 2);
  const stringFlagValue = client.stringVariation('example-flag-key', 'default');
  // etc.
});
```
