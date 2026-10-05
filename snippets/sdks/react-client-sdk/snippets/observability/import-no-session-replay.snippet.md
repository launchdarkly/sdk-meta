---
id: react-client-sdk/observability/import-no-session-replay
sdk: react-client-sdk
kind: import
lang: javascript
file: react-client-sdk/observability/import-no-session-replay.txt
description: Import statements for react-client-sdk with the observability plugin only (no session replay).
validation:
  scaffold: react-client-sdk/scaffolds/react-syntax-only
---

```javascript
import { withLDProvider } from 'launchdarkly-react-client-sdk';
import Observability, { LDObserve } from '@launchdarkly/observability'
```
