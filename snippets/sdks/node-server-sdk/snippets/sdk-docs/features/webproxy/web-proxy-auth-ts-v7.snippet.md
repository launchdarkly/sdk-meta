---
id: node-server-sdk/sdk-docs/features/webproxy/web-proxy-auth-ts-v7
sdk: node-server-sdk
kind: reference
lang: typescript
description: Web proxy configuration with authentication for Node.js (server-side, v7, TypeScript).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only
---

```typescript
import { LDOptions } from 'launchdarkly-node-server-sdk';

const options: LDOptions = {
  proxyHost: 'your-proxy-host',
  proxyPort: 8080,
  proxyAuth: 'username:password'
};
```
