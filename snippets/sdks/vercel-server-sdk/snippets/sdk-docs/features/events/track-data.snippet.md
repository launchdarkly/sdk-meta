---
id: vercel-server-sdk/sdk-docs/features/events/track-data
sdk: vercel-server-sdk
kind: reference
lang: typescript
description: Custom event with attached data and metric value for Vercel SDK v1.2.0+.
validation:
  scaffold: vercel-server-sdk/scaffolds/vercel-syntax-only

---

```typescript
const exampleData = { customProperty: 'someValue' };
const metricValue = 10;

client.track('example-event-key', context, exampleData, metricValue);
```
