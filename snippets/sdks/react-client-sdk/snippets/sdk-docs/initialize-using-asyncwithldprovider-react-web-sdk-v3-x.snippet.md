---
id: react-client-sdk/sdk-docs/initialize-using-asyncwithldprovider-react-web-sdk-v3-x
sdk: react-client-sdk
kind: reference
lang: javascript
description: "React Web SDK v3.x in section \"Initialize using `asyncWithLDProvider`\""
---

```js
import { asyncWithLDProvider } from 'launchdarkly-react-client-sdk';
import Observability from '@launchdarkly/observability';
import SessionReplay from '@launchdarkly/session-replay';

(async () => {
  const LDProvider = await asyncWithLDProvider({
    clientSideID: 'example-client-side-id',
    context: {
      "kind": "user",
      "key": "example-user-key",
      "name": "Sandy Smith",
      "email": "sandy@example.com"
    },
    options: {
      // the observability plugins require React Web SDK v3.7+
      plugins: [
        new Observability(),
        new SessionReplay()
      ],
      // other options...
    }
  });

  render(
    <LDProvider>
      <YourApp />
    </LDProvider>,
    document.getElementById('reactDiv'),
  );
})();
```
