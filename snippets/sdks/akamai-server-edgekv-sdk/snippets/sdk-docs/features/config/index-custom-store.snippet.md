---
id: akamai-server-edgekv-sdk/sdk-docs/features/config/index-custom-store
sdk: akamai-server-edgekv-sdk
kind: reference
lang: typescript
description: SDK configuration example for Akamai.
validation:
  scaffold: akamai-server-edgekv-sdk/scaffolds/akamai-syntax-only

---

```ts
// If you are using your own feature store,
// make sure to import 'init' from the 'akamai-server-base-sdk'
// rather than from 'akamai-server-edgekv-sdk'

import { init, EdgeProvider, LDOptions, BasicLogger } from '@launchdarkly/akamai-server-base-sdk';

// When using your own feature store, you must create a new class that
// implements EdgeProvider. To learn more, read
// https://launchdarkly.com/docs/sdk/edge/akamai#getting-started
class FeatureStore implements EdgeProvider {
  async get(rootKey: string): Promise<string> {
    return flagData;
  }
}

const options: LDOptions = {
  logger: BasicLogger.get(),
  cacheTtlMs: 200,
};

const ldClient = init({
  sdkKey: 'example-client-side-id',
  featureStoreProvider: new FeatureStore(),
  options: options,
});
```
