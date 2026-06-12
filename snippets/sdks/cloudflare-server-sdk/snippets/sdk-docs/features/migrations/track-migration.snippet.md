---
id: cloudflare-server-sdk/sdk-docs/features/migrations/track-migration
sdk: cloudflare-server-sdk
kind: reference
lang: typescript
description: Migration metrics tracking (trackMigration) for Cloudflare SDK v2.2.2+.
validation:
  scaffold: cloudflare-server-sdk/scaffolds/cloudflare-syntax-only
---

```typescript
import {
  LDClient,
  LDMigrationTracker,
} from '@launchdarkly/cloudflare-server-sdk';

const event = tracker.createEvent();

if (event) {
  client.trackMigration(event);
}
```
