---
id: fastly-server-sdk/sdk-docs/features/anonymous/anonymous
sdk: fastly-server-sdk
kind: reference
lang: typescript
description: Anonymous context example for Fastly.
validation:
  scaffold: fastly-server-sdk/scaffolds/edge-fastly-toplevel
---

```typescript
import type { LDContext } from '@launchdarkly/fastly-server-sdk';

const anonymousContext: LDContext = { kind: 'user', key: 'example-user-key', anonymous: true };
```
