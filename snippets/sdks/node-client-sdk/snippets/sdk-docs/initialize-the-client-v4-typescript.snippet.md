---
id: node-client-sdk/sdk-docs/initialize-the-client-v4-typescript
sdk: node-client-sdk
kind: reference
lang: typescript
description: "TypeScript, Node.js SDK v4.x in section \"Initialize the client\" (event listeners)"
validation:
  scaffold: node-client-sdk/scaffolds/node-client-syntax-only
---

```ts
client.on('ready', () => {
  // The client has finished starting up, whether or not it succeeded
});
client.on('initialized', () => {
  // Initialization succeeded, flag values are now available
  const boolFlagValue = client.variation('example-flag-key', false) as boolean;
  const numberFlagValue = client.variation('example-flag-key', 2) as number;
  const stringFlagValue = client.variation('example-flag-key', 'default') as string;
  // etc.
});
client.on('error', (context, err) => {
  // An error occurred, which may include an initialization failure
});
```
