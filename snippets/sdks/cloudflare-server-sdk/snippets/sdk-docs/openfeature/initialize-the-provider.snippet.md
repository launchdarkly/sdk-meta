---
id: cloudflare-server-sdk/sdk-docs/openfeature/initialize-the-provider
sdk: cloudflare-server-sdk
kind: reference
lang: javascript
file: cloudflare-server-sdk/sdk-docs/openfeature/initialize-the-provider.mjs
description: "Cloudflare OpenFeature provider in section \"Initialize the provider\""
validation:
  scaffold: cloudflare-server-sdk/scaffolds/cloudflare-syntax-only
---

```js
const provider = new LaunchDarklyProvider('example-client-side-id', env.LD_KV);
await OpenFeature.setProviderAndWait(provider);

const client = OpenFeature.getClient();
```
