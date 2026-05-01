---
id: react-client-sdk/scaffolds/react-syntax-only
sdk: react-client-sdk
kind: scaffold
lang: javascript
file: src/Snippet.jsx
description: |
  Parse-only validator for React client SDK doc fragments.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: react-client
  entrypoint: src/Snippet.jsx
---

```javascript
import React from 'react';

export function Snippet() {
{{ body }}
  return null;
}
```
