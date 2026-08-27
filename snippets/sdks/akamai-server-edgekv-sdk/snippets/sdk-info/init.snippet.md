---
id: akamai-server-edgekv-sdk/sdk-info/init
sdk: akamai-server-edgekv-sdk
kind: init
lang: typescript
file: akamai-server-edgekv-sdk/init.txt
description: |
  Client initialization snippet for akamai-server-edgekv-sdk.
  Type-checked against the real package via the edge-akamai-toplevel
  scaffold (runtime validation would need a provisioned Akamai EdgeKV).
  Note: the Akamai SDK has no waitForInitialization — every variation
  call reads from EdgeKV.
validation:
  scaffold: akamai-server-edgekv-sdk/scaffolds/edge-akamai-toplevel
---

```typescript
import { init } from '@launchdarkly/akamai-server-edgekv-sdk';

// Initialize the client with your LaunchDarkly client-side ID and the
// EdgeKV namespace and group configured for the LaunchDarkly Akamai
// integration.
const ldClient = init({
  sdkKey: 'YOUR_CLIENT_SIDE_ID',
  namespace: 'your-edgekv-namespace',
  group: 'your-edgekv-group-id',
});
```
