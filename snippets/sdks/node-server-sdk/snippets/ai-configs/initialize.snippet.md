---
id: node-server-sdk/ai-configs/initialize
sdk: node-server-sdk
kind: initialize
lang: javascript
file: node-server-sdk/ai-configs/initialize.txt
description: Initialize the LaunchDarkly client and AI Configs client for node-server-sdk.
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only
---

```javascript
const ldClient: LDClient = init('{{sdkkey}}');
try {
    await ldClient.waitForInitialization({ timeout: 10 });
    // initialization complete
} catch (error) {
    // timeout or SDK failed to initialize
}
const aiClient: LDAIClient = initAi(ldClient);
```
