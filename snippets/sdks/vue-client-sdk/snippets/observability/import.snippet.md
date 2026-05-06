---
id: vue-client-sdk/observability/import
sdk: vue-client-sdk
kind: import
lang: javascript
file: vue-client-sdk/observability/import.txt
description: Import statements for vue-client-sdk observability + session replay plugins.
---

```javascript
import { createApp } from 'vue'
import App from './App.vue'
import { LDPlugin } from 'launchdarkly-vue-client-sdk'
import Observability, { LDObserve } from '@launchdarkly/observability'
import SessionReplay, { LDRecord } from '@launchdarkly/session-replay'
```
