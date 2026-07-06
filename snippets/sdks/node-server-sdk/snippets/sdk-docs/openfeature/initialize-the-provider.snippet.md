---
id: node-server-sdk/sdk-docs/openfeature/initialize-the-provider
sdk: node-server-sdk
kind: reference
lang: javascript
file: node-server-sdk/sdk-docs/openfeature/initialize-the-provider.mjs
description: "Node.js OpenFeature provider in section \"Initialize the provider\""
validation:
  scaffold: node-server-sdk/scaffolds/openfeature-node-init-runner
  placeholders:
    YOUR_SDK_KEY: LAUNCHDARKLY_SDK_KEY
---

```js
await OpenFeature.setProviderAndWait(new LaunchDarklyProvider("YOUR_SDK_KEY"));

const client = OpenFeature.getClient();
```
