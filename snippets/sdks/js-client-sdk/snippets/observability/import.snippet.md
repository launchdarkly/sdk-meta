---
id: js-client-sdk/observability/import
sdk: js-client-sdk
kind: import
lang: javascript
file: js-client-sdk/observability/import.txt
description: Import statements for js-client-sdk observability + session replay plugins.
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only
---

```javascript
import LDClient from 'launchdarkly-js-client-sdk'
import Observability, { LDObserve } from '@launchdarkly/observability'
import SessionReplay, { LDRecord } from '@launchdarkly/session-replay'
```
