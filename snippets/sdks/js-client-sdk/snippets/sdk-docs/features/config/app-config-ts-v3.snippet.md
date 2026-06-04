---
id: js-client-sdk/sdk-docs/features/config/app-config-ts-v3
sdk: js-client-sdk
kind: reference
lang: typescript
description: Application metadata configuration example for JavaScript.
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only

---

```ts
import * as LDClient from 'launchdarkly-js-client-sdk';

const options: LDClient.LDOptions = {
    application: {
	    id: 'authentication-service',
	    version: '1.0.0',
    },
};
const client = LDClient.initialize('example-client-side-id', context, options);
```
