---
id: node-server-sdk/sdk-docs/features/otel/resource-attributes
sdk: node-server-sdk
kind: reference
lang: javascript
description: Programmatic OpenTelemetry resource attribute configuration (Node.js, server-side).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only
---

```javascript
import { Resource } from '@opentelemetry/resources';
import { NodeSDK } from '@opentelemetry/sdk-node';

const sdk = new NodeSDK({
  resource: new Resource({
    'launchdarkly.project_id': 'YOUR_SDK_KEY',
  }),
});

sdk.start();
```
