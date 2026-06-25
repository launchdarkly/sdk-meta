---
id: node-server-sdk/sdk-docs/features/webproxy/web-proxy-auth-ts-v8
sdk: node-server-sdk
kind: reference
lang: typescript
description: Web proxy configuration with authentication for Node.js (server-side, v8, TypeScript).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only
---

```typescript
import { LDOptions } from '@launchdarkly/node-server-sdk';

const options: LDOptions = {
  proxyOptions: {
      host: 'your-proxy-host',
      port: 8080,
      scheme: 'https',
      auth: 'username:password'
    }
};
```
