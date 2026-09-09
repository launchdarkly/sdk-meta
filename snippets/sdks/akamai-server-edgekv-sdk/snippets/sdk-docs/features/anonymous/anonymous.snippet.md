---
id: akamai-server-edgekv-sdk/sdk-docs/features/anonymous/anonymous
sdk: akamai-server-edgekv-sdk
kind: reference
lang: typescript
description: Anonymous context example for Akamai.
validation:
  scaffold: akamai-server-edgekv-sdk/scaffolds/edge-akamai-toplevel
---

```typescript
import type { LDContext } from '@launchdarkly/akamai-server-edgekv-sdk';

const anonymousContext: LDContext = { kind: 'user', key: 'example-user-key', anonymous: true };
```
