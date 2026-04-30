---
id: js-client-sdk/getting-started/app-ts
sdk: js-client-sdk
kind: hello-world
lang: ts
file: src/app.ts
description: Browser app entry that initializes the JavaScript client SDK and renders the flag value with initialization status.
inputs:
  environmentId:
    type: client-side-id
    description: Client-side ID baked into the rendered source. Validation reads LAUNCHDARKLY_CLIENT_SIDE_ID at runtime.
  featureKey:
    type: flag-key
    description: Default flag key baked into the rendered source. Validation reads LAUNCHDARKLY_FLAG_KEY at runtime.
validation:
  runtime: js-client
  entrypoint: src/app.ts
  companions: [js-client-sdk/getting-started/index-html]
---

Initialize the SDK and render the flag value in `src/app.ts`:

```ts
import { createClient } from '@launchdarkly/js-client-sdk';

const clientSideID = '{{ environmentId }}';
const flagKey = '{{ featureKey }}';

const context = {
  kind: 'user',
  key: 'example-user-key',
  name: 'Sandy',
};

const STATUS_MESSAGES = {
  complete: 'SDK successfully initialized!',
  failed: 'SDK failed to initialize. Please check your internet connection and SDK credential for any typo.',
  timeout: 'Timeout identifying client',
} as const;

const statusBox = document.createElement('div');
const flagBox = document.createElement('div');
document.body.append(statusBox, flagBox);
statusBox.textContent = 'Initializing…';

const main = async () => {
  const ldclient = createClient(clientSideID, context);

  const render = () => {
    const flagValue = ldclient.variation(flagKey, false);
    document.body.style.background = flagValue ? '#00844B' : '#373841';
    flagBox.textContent = `The ${flagKey} feature flag evaluates to ${flagValue}.`;
  };

  ldclient.on('error', () => { statusBox.textContent = 'Error caught in client SDK'; });
  ldclient.on('change', render);

  ldclient.start();
  const { status } = await ldclient.waitForInitialization();
  statusBox.textContent = STATUS_MESSAGES[status] ?? 'Unknown error identifying client';
  render();
};

main();
```
