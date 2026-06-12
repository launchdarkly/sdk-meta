---
id: node-server-sdk/sdk-docs/features/aimetrics/get-summary-agent
sdk: node-server-sdk
kind: reference
lang: typescript
description: Retrieve automatically recorded metrics with getSummary in agent mode for the Node.js (server-side) AI SDK.
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```typescript
const summary = agent.tracker.getSummary();

// recorded metrics available in summary.durationMs, summary.feedback,
// summary.tokens, summary.success, summary.timeToFirstTokenMs

```
