---
id: vercel-server-sdk/sdk-docs/features/migrations/track-migration
sdk: vercel-server-sdk
kind: reference
lang: typescript
description: Migration metrics tracking (trackMigration) for Vercel SDK v1.1.6+.
validation:
  scaffold: vercel-server-sdk/scaffolds/vercel-syntax-only
---

```typescript
import {
  LDMigrationTracker,
  LDClient
} from '@launchdarkly/vercel-server-sdk';

const event = tracker.createEvent();

if (event) {
  client.trackMigration(event);
}
```
