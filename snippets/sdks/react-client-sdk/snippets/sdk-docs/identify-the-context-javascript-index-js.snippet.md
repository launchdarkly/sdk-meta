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
import ReactDOM from 'react-dom';
import App from 'app.js';
import { asyncWithLDProvider } from 'launchdarkly-react-client-sdk';

const renderApp = async () => {
  const LDProvider = await asyncWithLDProvider({ clientSideID: 'example-client-side-id' });

  ReactDOM.render(
    <React.StrictMode>
      <LDProvider>
        <App />
      </LDProvider>
    </React.StrictMode>,
    document.getElementById('root')
  )
}

renderApp();
```
