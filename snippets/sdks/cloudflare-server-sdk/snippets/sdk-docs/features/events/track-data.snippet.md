---
id: cloudflare-server-sdk/sdk-docs/features/events/track-data
sdk: cloudflare-server-sdk
kind: reference
lang: typescript
description: Custom event with attached data and metric value for Cloudflare SDK v2.3.0+.
validation:
  scaffold: cloudflare-server-sdk/scaffolds/cloudflare-syntax-only

---

```typescript
const exampleData = { customProperty: 'someValue' };
const metricValue = 10;

client.track('example-event-key', context, exampleData, metricValue);
```
