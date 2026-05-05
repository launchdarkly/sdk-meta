---
id: js-client-sdk/sdk-info/init
sdk: js-client-sdk
kind: init
lang: javascript
file: js-client-sdk/init.txt
description: Client initialization snippet for js-client-sdk.
validation:
  scaffold: js-client-sdk/scaffolds/init-runner
  placeholders:
    YOUR_CLIENT_SIDE_ID: LAUNCHDARKLY_CLIENT_SIDE_ID
---

```javascript
// A "context" is a data object representing users, devices, organizations, and
// other entities. You'll need this later, but you can ignore it for now.
const context = {
  kind: 'user',
  key: 'EXAMPLE_CONTEXT_KEY'
};
// This is your client-side ID.
const client = createClient('YOUR_CLIENT_SIDE_ID', context);
client.start();

const { status } = await client.waitForInitialization();

if (status === 'complete') {
  console.log('SDK successfully initialized!');
} else {
  console.error('Initialization failed');
}
```
