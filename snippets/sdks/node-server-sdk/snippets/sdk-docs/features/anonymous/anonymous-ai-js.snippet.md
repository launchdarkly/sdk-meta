---
id: node-server-sdk/sdk-docs/features/anonymous/anonymous-ai-js
sdk: node-server-sdk
kind: reference
lang: javascript
description: Anonymous context example for the Node.js (server-side) AI SDK (JavaScript).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only
---

```javascript
const context = {
  kind: 'user',
  key: 'example-user-key',
  anonymous: true,
}
```
