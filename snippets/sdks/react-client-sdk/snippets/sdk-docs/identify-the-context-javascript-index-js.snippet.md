---
id: react-client-sdk/sdk-docs/identify-the-context-javascript-index-js
sdk: react-client-sdk
kind: reference
lang: javascript
description: "JavaScript: index.js in section \"Identify the context\""
---

```js
// index.js
import React from 'react';
import ReactDOM from 'react-dom/client';
import App from 'app.js';
import { createLDReactProvider } from '@launchdarkly/react-sdk';

const root = ReactDOM.createRoot(document.getElementById('root'));

const LDProvider = createLDReactProvider('example-client-side-id', {
  kind: 'user',
  anonymous: true,
});

root.render(
  <LDProvider>
    <App />
  </LDProvider>
);
```
