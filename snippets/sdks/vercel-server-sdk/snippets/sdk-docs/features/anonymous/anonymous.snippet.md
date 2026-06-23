---
id: vercel-server-sdk/sdk-docs/features/anonymous/anonymous
sdk: vercel-server-sdk
kind: reference
lang: typescript
description: Anonymous context example for Vercel.
validation:
  scaffold: vercel-server-sdk/scaffolds/vercel-syntax-only
---

```typescript
import type { LDContext } from '@launchdarkly/vercel-server-sdk';

const anonymousContext: LDContext = { kind: 'user', key: 'example-user-key', anonymous: true };
```
