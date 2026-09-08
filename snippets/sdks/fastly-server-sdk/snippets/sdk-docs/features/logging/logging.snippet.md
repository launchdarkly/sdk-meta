---
id: fastly-server-sdk/sdk-docs/features/logging/logging
sdk: fastly-server-sdk
kind: reference
lang: typescript
description: basicLogger debug-level configuration example for Fastly.
validation:
  scaffold: fastly-server-sdk/scaffolds/edge-fastly-toplevel

---

```typescript
import { basicLogger, LDOptions } from '@launchdarkly/fastly-server-sdk';

const options: LDOptions = {
  logger: basicLogger({ level: 'debug' }),
};
```
