---
id: react-client-sdk/sdk-docs/multi-environment-support-environments-ts
sdk: react-client-sdk
kind: reference
lang: javascript
description: "JavaScript: environments.ts in section \"Multi-environment support\""
---

```js
// Single module, import anywhere in your app
import {
  initLDReactContext,
  createClient,
  createLDReactProviderWithClient,
} from '@launchdarkly/react-sdk';

export const ProdLDContext    = initLDReactContext();
export const StagingLDContext = initLDReactContext();

const prodClient    = createClient('prod-client-side-id',    { kind: 'user', key: 'user-key' });
const stagingClient = createClient('staging-client-side-id', { kind: 'user', key: 'user-key' });
prodClient.start();
stagingClient.start();

export const ProdLDProvider    = createLDReactProviderWithClient(prodClient,    ProdLDContext);
export const StagingLDProvider = createLDReactProviderWithClient(stagingClient, StagingLDContext);
```
