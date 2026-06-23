---
id: electron-client-sdk/sdk-docs/features/anonymous/anonymous-ts
sdk: electron-client-sdk
kind: reference
lang: typescript
description: Anonymous user example for Electron (TypeScript).
validation:
  scaffold: electron-client-sdk/scaffolds/electron-syntax-only
---

```typescript
import * as LDElectron from 'launchdarkly-electron-client-sdk';

const anonymousUser: LDElectron.LDUser = { key: 'example-user-key', anonymous: true };
```
