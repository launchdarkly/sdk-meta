---
id: js-client-sdk/sdk-docs/features/hooks/add-hook-v4
sdk: js-client-sdk
kind: reference
lang: js
description: Adding a hook to an existing client for the JavaScript SDK v4.x.
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only
---

```js

const client = LDClient.initialize('example-client-side-id', options);
client.addHook(new ExampleHook());

```
