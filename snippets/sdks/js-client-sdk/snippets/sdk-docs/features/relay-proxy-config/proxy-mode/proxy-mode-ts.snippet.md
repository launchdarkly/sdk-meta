---
id: js-client-sdk/sdk-docs/features/relay-proxy-config/proxy-mode/proxy-mode-ts
sdk: js-client-sdk
kind: reference
lang: typescript
description: Proxy mode configuration example for JavaScript.
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only
---

```ts
import * as ld from 'launchdarkly-js-client-sdk';

const options: ld.LDOptions = {
  streamUrl: 'https://your-relay-proxy.com:8030',
  baseUrl: 'https://your-relay-proxy.com:8030',
  eventsUrl: 'https://your-relay-proxy.com:8030',
};
```
