---
id: node-server-sdk/sdk-docs/features/aimetrics/agent-instructions
sdk: node-server-sdk
kind: reference
lang: javascript
description: Access instructions and record metrics in agent mode for the Node.js (server-side) AI SDK.
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```javascript
if (agent.enabled) {

  // Retrieve instructions from the config and pass to your AI model
  const result = example_model_api(agent.instructions)

  // Track metrics from the result
  agent.tracker.trackSuccess()

} else {

  // Application path to take when the agent config is disabled

}
```
