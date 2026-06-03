---
id: react-client-sdk/sdk-docs/createldserversession-and-useldserversession-use-a-server-session
sdk: react-client-sdk
kind: reference
lang: javascript
description: "JavaScript: Use a server session in section \"`createLDServerSession` and `useLDServerSession`\""
validation:
  scaffold: react-client-sdk/scaffolds/react-syntax-only
---

```js
// components/FeatureBanner.tsx — retrieve it in a nested component
import { useLDServerSession } from '@launchdarkly/react-sdk/server';

export default async function FeatureBanner() {
  const session = useLDServerSession();
  if (!session) return null;

  const banner = await session.stringVariation('banner-text', 'Welcome');
  return <h1>{banner}</h1>;
}
```
