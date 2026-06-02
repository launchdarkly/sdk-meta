---
id: js-client-sdk/sdk-docs/features/config/index-ts-v3
sdk: js-client-sdk
kind: reference
lang: typescript
description: SDK configuration example for JavaScript.
---

```ts
import * as LDClient from 'launchdarkly-js-client-sdk';

const options: LDClient.LDOptions = { allAttributesPrivate: true };
const client = LDClient.initialize('example-client-side-id', context, options);
```
