---
id: vercel-server-sdk/sdk-docs/features/logging/logging
sdk: vercel-server-sdk
kind: reference
lang: typescript
description: BasicLogger debug-level configuration example for Vercel.
validation:
  scaffold: vercel-server-sdk/scaffolds/vercel-syntax-only

---

```typescript
import { BasicLogger, LDOptions } from '@launchdarkly/vercel-server-sdk';

const options: LDOptions = {
  logger: new BasicLogger({ level: 'debug', }),
};
```
