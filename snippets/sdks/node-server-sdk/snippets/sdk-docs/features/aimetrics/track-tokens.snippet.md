---
id: node-server-sdk/sdk-docs/features/aimetrics/track-tokens
sdk: node-server-sdk
kind: reference
lang: typescript
description: Track token usage manually for the Node.js (server-side) AI SDK.
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```typescript
import { LDTokenUsage } from '@launchdarkly/server-sdk-ai';

// Track your own token usage.

// First, set up an LDTokenUsage object.
// Update the input, output, and total fields
// with return values from your AI model generation.
const tokens: LDTokenUsage = {
  input: 0,
  output: 0,
  total: 0,
}

aiConfig.tracker.trackTokens(tokens);
```
