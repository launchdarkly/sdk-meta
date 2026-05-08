---
id: react-client-sdk/observability/import
sdk: react-client-sdk
kind: import
lang: javascript
file: react-client-sdk/observability/import.txt
description: Import statements for react-client-sdk observability + session replay plugins.
validation:
  scaffold: react-client-sdk/scaffolds/react-syntax-only
---

```javascript
import { withLDProvider } from 'launchdarkly-react-client-sdk';
import Observability, { LDObserve } from '@launchdarkly/observability'
import SessionReplay, { LDRecord } from '@launchdarkly/session-replay'
```
