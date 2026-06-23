---
id: cloudflare-server-sdk/sdk-docs/features/logging/logging
sdk: cloudflare-server-sdk
kind: reference
lang: typescript
description: BasicLogger debug-level configuration example for Cloudflare.
validation:
  scaffold: cloudflare-server-sdk/scaffolds/cloudflare-syntax-only

---

```typescript
import { BasicLogger, LDOptions } from '@launchdarkly/cloudflare-server-sdk';

const options: LDOptions = {
  logger: new BasicLogger({ level: 'debug', }),
};
```
