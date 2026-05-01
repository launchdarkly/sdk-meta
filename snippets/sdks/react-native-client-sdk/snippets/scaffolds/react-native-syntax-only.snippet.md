---
id: react-native-client-sdk/scaffolds/react-native-syntax-only
sdk: react-native-client-sdk
kind: scaffold
lang: javascript
file: App.js
description: |
  Parse-only validator for React Native client SDK doc fragments.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: react-native-client
  entrypoint: App.js
---

```javascript
import React from 'react';

export default function App() {
{{ body }}
  return null;
}
```
