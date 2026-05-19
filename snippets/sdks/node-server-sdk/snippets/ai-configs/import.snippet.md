---
id: node-server-sdk/ai-configs/import
sdk: node-server-sdk
kind: import
lang: javascript
file: node-server-sdk/ai-configs/import.txt
description: Import statements for node-server-sdk AI Configs.
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only
---

```javascript
import { init, LDClient } from '@launchdarkly/node-server-sdk';
import { initAi, LDAIClient, LDAIConfig } from '@launchdarkly/server-sdk-ai';
```
