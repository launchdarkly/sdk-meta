---
id: node-server-sdk/ai-configs/initialize
sdk: node-server-sdk
kind: initialize
lang: javascript
file: node-server-sdk/ai-configs/initialize.txt
description: Initialize the LaunchDarkly client and AI Configs client for node-server-sdk.
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
