---
id: react-client-sdk/sdk-docs/features/datasaving/fixed-connection-mode
sdk: react-client-sdk
kind: reference
lang: javascript
description: Set a fixed connection mode with manual mode switching for React Web.
validation:
  scaffold: react-client-sdk/scaffolds/react-syntax-only
---

```js
export const LDReactProvider = createLDReactProvider('example-client-side-id', context, {
  ldOptions: {
    dataSystem: { automaticModeSwitching: { type: 'manual', initialConnectionMode: 'polling' } },
  },
});
```
