---
id: js-client-sdk/sdk-docs/use-events-to-determine-when-the-client-is-ready-javascript-sdk-v3-x-v4-x
sdk: js-client-sdk
kind: reference
lang: javascript
description: "JavaScript SDK, v3.x, v4.x in section \"Use events to determine when the client is ready\""
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only
---

```js
client.on('ready', () => {
  // initialization succeeded, flag values are now available
  const flagValue = client.variation('example-flag-key', false);
  // etc.
});
```
