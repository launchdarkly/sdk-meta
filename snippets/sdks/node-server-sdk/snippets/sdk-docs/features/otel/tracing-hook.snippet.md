---
id: node-server-sdk/sdk-docs/features/otel/tracing-hook
sdk: node-server-sdk
kind: reference
lang: ts
description: OpenTelemetry tracing hook configuration for the Node.js (server-side) SDK.
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only
---

```ts
import { LDOptions, init } from '@launchdarkly/node-server-sdk';
import { TracingHook } from '@launchdarkly/node-server-sdk-otel';

const options: LDOptions = {
  hooks: [new TracingHook()]
}

const client = init('YOUR_SDK_KEY', options);

```
