---
id: node-server-sdk/sdk-docs/promises-and-async-node-js-sdk-v7-x-and-later-typescript
sdk: node-server-sdk
kind: reference
lang: typescript
description: "Node.js SDK v7.x and later (TypeScript) in section \"Promises and async\""
---

```ts
// Using the .then() method to add a continuation handler for a Promise
client.variation('example-flag-key', context, false).then((value) => {
  // application code
});

// Using "await" instead, within an async function
const value = await client.variation('example-flag-key', context, false);

// In both cases, you can cast "value" to a boolean, number, or string,
// rather than using the LDFlagValue type,
// if you know the type of your flag variations
```
