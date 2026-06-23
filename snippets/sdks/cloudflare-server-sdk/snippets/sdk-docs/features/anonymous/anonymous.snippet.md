---
id: cloudflare-server-sdk/sdk-docs/features/anonymous/anonymous
sdk: cloudflare-server-sdk
kind: reference
lang: typescript
description: Anonymous context example for Cloudflare.
validation:
  scaffold: cloudflare-server-sdk/scaffolds/cloudflare-syntax-only
---

```typescript
import type { LDContext } from '@launchdarkly/cloudflare-server-sdk';

const anonymousContext: LDContext = { kind: 'user', key: 'example-user-key', anonymous: true };
```
