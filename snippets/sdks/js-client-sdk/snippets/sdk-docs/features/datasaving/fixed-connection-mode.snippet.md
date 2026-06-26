---
id: js-client-sdk/sdk-docs/features/datasaving/fixed-connection-mode
sdk: js-client-sdk
kind: reference
lang: javascript
description: Set a fixed connection mode with manual mode switching for JavaScript.
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only
---

```js
const client = createClient('example-client-side-id', context, {
  dataSystem: {
    automaticModeSwitching: { type: 'manual', initialConnectionMode: 'polling' },
  },
});
```
