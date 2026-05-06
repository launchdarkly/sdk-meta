---
id: react-client-sdk/sdk-docs/multi-environment-support-in-any-component
sdk: react-client-sdk
kind: reference
lang: javascript
description: "JavaScript: In any component in section \"Multi-environment support\""
validation:
  scaffold: react-client-sdk/scaffolds/react-syntax-only
---

```js
import { useBoolVariation } from '@launchdarkly/react-sdk';
import { ProdLDContext, StagingLDContext } from './environments';

function MyComponent() {
  const showInProd    = useBoolVariation('my-feature', false, ProdLDContext);
  const showInStaging = useBoolVariation('my-feature', false, StagingLDContext);
  // ...
}
```
