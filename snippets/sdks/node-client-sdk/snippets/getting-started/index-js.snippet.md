---
id: node-client-sdk/getting-started/index-js
sdk: node-client-sdk
kind: hello-world
lang: javascript
file: index.js
description: Hello-world program that initializes the Node.js client SDK and evaluates a feature flag.
inputs:
  environmentId:
    type: client-side-id
    description: Client-side ID baked into the rendered source. Validation reads LAUNCHDARKLY_CLIENT_SIDE_ID at runtime.
  featureKey:
    type: flag-key
    description: Default flag key baked into the rendered source. Validation reads LAUNCHDARKLY_FLAG_KEY at runtime.
ld-application:
  slot: index-js
validation:
  runtime: node
  requirements: launchdarkly-node-client-sdk
---

Create a file called `index.js` and add the following code:

```javascript
// Import the LaunchDarkly client
var LaunchDarkly = require('launchdarkly-node-client-sdk');

// Set up the user properties. This user should appear on your LaunchDarkly users dashboard
// soon after you run the demo.
var user = {
  key: "example-user-key"
};

// Create a single instance of the LaunchDarkly client
const ldClient = LaunchDarkly.initialize('{{ environmentId }}', user);

function showMessage(s) {
  console.log("*** " + s);
  console.log("");
}
ldClient.waitForInitialization().then(function() {
  showMessage("SDK successfully initialized!");
  const flagValue = ldClient.variation("{{ featureKey }}", false);

  showMessage("The '" + "{{ featureKey }}" + "' feature flag evaluates to " + flagValue + ".");

  // Here we ensure that the SDK shuts down cleanly and has a chance to deliver analytics
  // events to LaunchDarkly before the program exits. If analytics events are not delivered,
  // the user properties and flag usage statistics will not appear on your dashboard. In a
  // normal long-running application, the SDK would continue running and events would be
  // delivered automatically in the background.
  ldClient.close();
}).catch(function(error) {
  showMessage("SDK failed to initialize: " + error);
  process.exit(1);
});
```
