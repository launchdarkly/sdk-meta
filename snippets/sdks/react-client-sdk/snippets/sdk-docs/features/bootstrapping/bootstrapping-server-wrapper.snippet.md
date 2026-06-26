---
id: react-client-sdk/sdk-docs/features/bootstrapping/bootstrapping-server-wrapper
sdk: react-client-sdk
kind: reference
lang: javascript
description: Bootstrapping through the React Server Wrapper (LDIsomorphicProvider).
validation:
  scaffold: react-client-sdk/scaffolds/react-server-syntax-only
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
