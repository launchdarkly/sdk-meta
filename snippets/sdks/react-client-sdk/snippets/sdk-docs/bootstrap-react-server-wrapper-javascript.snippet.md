---
id: react-client-sdk/sdk-docs/bootstrap-react-server-wrapper-javascript
sdk: react-client-sdk
kind: reference
lang: javascript
description: "JavaScript in section \"React Server Wrapper\""
# TODO(scaffold): React Server Component snippet — imports
# @launchdarkly/node-server-sdk + @launchdarkly/react-sdk/server,
# neither of which the react-client validator container pre-bakes.
# Needs a dedicated react-server-syntax-only scaffold paired with a
# react-server validator (or react-client deps extended with
# node-server-sdk + react-sdk/server subpath). Once that scaffold
# lands, set validation.scaffold here.
---

```js
// app/page.tsx (Server Component)
import { init } from '@launchdarkly/node-server-sdk';
import { createLDServerSession, LDIsomorphicProvider } from '@launchdarkly/react-sdk/server';

const ldBaseClient = init(process.env.LAUNCHDARKLY_SDK_KEY!);

export default async function Page() {
  await ldBaseClient.waitForInitialization({ timeout: 10 });

  const session = createLDServerSession(ldBaseClient, {
    kind: 'user',
    key: 'user-key',
    name: 'Sandy',
  });

  return (
    <LDIsomorphicProvider
      session={session}
      clientSideId={process.env.LAUNCHDARKLY_CLIENT_SIDE_ID!}
    >
      <App />
    </LDIsomorphicProvider>
  );
}
```
