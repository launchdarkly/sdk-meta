---
id: node-server-sdk/sdk-docs/features/webproxy/web-proxy-auth-js-v8
sdk: node-server-sdk
kind: reference
lang: javascript
description: Web proxy configuration with authentication for Node.js (server-side, v8, JavaScript).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only
---

```javascript

const options = {
  proxyOptions: {
      host: 'your-proxy-host',
      port: 8080,
      scheme: 'https',
      auth: 'username:password'
    }
};
```
