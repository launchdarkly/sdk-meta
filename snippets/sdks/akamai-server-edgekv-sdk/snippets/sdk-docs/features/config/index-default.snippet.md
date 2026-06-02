---
id: akamai-server-edgekv-sdk/sdk-docs/features/config/index-default
sdk: akamai-server-edgekv-sdk
kind: reference
lang: typescript
description: SDK configuration example for Akamai.
---

```ts
import { init, LDOptions, BasicLogger } from '@launchdarkly/akamai-server-edgekv-sdk';

const options: LDOptions = {
  logger: BasicLogger.get(),
  cacheTtlMs: 200,
};

const ldClient = init({
  sdkKey: 'example-client-side-id',
  namespace: 'your-edgekv-namespace',
  group: 'your-edgekv-group-id',
  options: options,
});
```
