---
id: react-client-sdk/sdk-docs/react-server-component-support-create-a-wrapped-server-side-sdk-client-react-web-sdk-v4-0
sdk: react-client-sdk
kind: reference
lang: tsx
description: "React Web SDK v4.0 in section \"Create a wrapped server-side SDK client\""
validation:
  scaffold: react-client-sdk/scaffolds/react-syntax-only
---

```tsx
import { init } from '@launchdarkly/node-server-sdk';
import { createLDServerSession } from '@launchdarkly/react-sdk/server';

const ldBaseClient = init(process.env.LAUNCHDARKLY_SDK_KEY!);

// In your root Server Component (for example, app/layout.tsx or app/page.tsx)
export default async function Page() {
  await ldBaseClient.waitForInitialization({ timeout: 10 });

  const session = createLDServerSession(ldBaseClient, {
    kind: 'user',
    key: 'user-key',
    name: 'Sandy',
  });

  const showFeature = await session.boolVariation('show-new-feature', false);
  return <div>{showFeature ? 'New Feature!' : 'Classic'}</div>;
}
```
