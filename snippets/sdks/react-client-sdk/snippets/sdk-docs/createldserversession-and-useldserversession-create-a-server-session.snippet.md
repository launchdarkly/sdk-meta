---
id: react-client-sdk/sdk-docs/createldserversession-and-useldserversession-create-a-server-session
sdk: react-client-sdk
kind: reference
lang: javascript
description: "JavaScript: Create a server session in section \"`createLDServerSession` and `useLDServerSession`\""
---

```js
// app/page.tsx — create the session in a parent component
import { createLDServerSession } from '@launchdarkly/react-sdk/server';

export default async function Page() {
  await ldBaseClient.waitForInitialization({ timeout: 10 });
  createLDServerSession(ldBaseClient, { kind: 'user', key: 'user-key' });

  return <FeatureBanner />;
}
```
