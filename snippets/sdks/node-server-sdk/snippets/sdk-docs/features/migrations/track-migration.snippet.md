---
id: node-server-sdk/sdk-docs/features/migrations/track-migration
sdk: node-server-sdk
kind: reference
lang: typescript
description: Migration metrics tracking (trackMigration) for Node.js (server-side) SDK v9.
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only
---

```typescript
const event = tracker.createEvent();

if (event) {
  client.trackMigration(event);
}
```
