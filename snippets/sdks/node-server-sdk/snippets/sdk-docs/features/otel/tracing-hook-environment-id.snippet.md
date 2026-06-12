---
id: node-server-sdk/sdk-docs/features/otel/tracing-hook-environment-id
sdk: node-server-sdk
kind: reference
lang: ts
description: OpenTelemetry tracing hook with an explicit environment ID for the Node.js (server-side) SDK.
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only
---

```ts
import { LDOptions, init } from '@launchdarkly/node-server-sdk';
import { TracingHook } from '@launchdarkly/node-server-sdk-otel'; // v1.2 or later

const options: LDOptions = {
  hooks: [new TracingHook({
    spans: true,
    includeValue: true,
    environmentId: 'example-client-side-id'
  })]
}

const client = init('YOUR_SDK_KEY', options);
```
