---
id: fastly-server-sdk/sdk-docs/features/logging/logging
sdk: fastly-server-sdk
kind: reference
lang: typescript
description: BasicLogger debug-level configuration example for Fastly.
validation:
  scaffold: fastly-server-sdk/scaffolds/fastly-syntax-only

---

```typescript
import { BasicLogger, LDOptions } from '@launchdarkly/fastly-server-sdk';

const options: LDOptions = {
  logger: new BasicLogger({ level: 'debug', }),
};
```
