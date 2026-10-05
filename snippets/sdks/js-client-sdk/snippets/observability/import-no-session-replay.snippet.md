---
id: js-client-sdk/observability/import-no-session-replay
sdk: js-client-sdk
kind: import
lang: javascript
file: js-client-sdk/observability/import-no-session-replay.txt
description: Import statements for js-client-sdk with the observability plugin only (no session replay).
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only
---

```javascript
import { createClient } from '@launchdarkly/js-client-sdk'
import Observability, { LDObserve } from '@launchdarkly/observability'
```
