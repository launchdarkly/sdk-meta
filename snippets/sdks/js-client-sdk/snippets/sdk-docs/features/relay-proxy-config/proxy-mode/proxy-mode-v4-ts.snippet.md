---
id: js-client-sdk/sdk-docs/features/relay-proxy-config/proxy-mode/proxy-mode-v4-ts
sdk: js-client-sdk
kind: reference
lang: ts
description: Proxy mode configuration example for JavaScript SDK v4.x (TypeScript).
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only
---

```ts
import type { LDOptions } from '@launchdarkly/js-client-sdk';

const options: LDOptions = {
  streamUri: 'https://your-relay-proxy.com:8030',
  baseUri: 'https://your-relay-proxy.com:8030',
  eventsUri: 'https://your-relay-proxy.com:8030',
};
```
