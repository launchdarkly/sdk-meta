---
id: js-client-sdk/sdk-docs/use-promises-to-determine-when-the-client-is-ready-javascript-sdk-v4-x-typescript
sdk: js-client-sdk
kind: reference
lang: typescript
description: "JavaScript SDK, v4.x (TypeScript) in section \"Use promises to determine when the client is ready\""
---

```ts
  import { initialize, LDClient } from '@launchdarkly/js-client-sdk';
  import { Observability } from '@launchdarkly/observability';
  import { SessionReplay } from '@launchdarkly/session-replay';

  const client: LDClient = initialize('example-client-side-id', {
    plugins: [
      new Observability(),
      new SessionReplay()
    ]
  });

  try {
    await client.waitForInitialization(5);
    // initialization succeeded, flag values are now available
    handleInitializedClient(client);
  } catch (err) {
    // initialization failed or did not complete before timeout
  }

  // Example user-defined function
  function handleInitializedClient(client: LDClient) {
    // Replace this with your application's logic
    const boolFlagValue = client.variation('example-flag-key', false) as boolean;
    console.log("LaunchDarkly client initialized successfully", boolFlagValue);
  }
```
