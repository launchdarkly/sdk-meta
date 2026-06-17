---
id: fastly-server-sdk/sdk-docs/features/events/track-data
sdk: fastly-server-sdk
kind: reference
lang: typescript
description: Custom event with attached data and metric value for Fastly.
validation:
  scaffold: fastly-server-sdk/scaffolds/fastly-syntax-only

---

```typescript
const exampleData = { customProperty: 'someValue' };
const metricValue = 10;

client.track('example-event-key', context, exampleData, metricValue);
```
