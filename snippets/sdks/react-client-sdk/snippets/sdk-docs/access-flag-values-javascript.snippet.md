---
id: react-client-sdk/sdk-docs/access-flag-values-javascript
sdk: react-client-sdk
kind: reference
lang: javascript
description: "JavaScript in section \"Access flag values\""
---

```js
import { useBoolVariation } from '@launchdarkly/react-sdk';

const Home = () => {
  const showNewFeature = useBoolVariation('show-new-feature', false);

  return showNewFeature ? <div>Flag on</div> : <div>Flag off</div>;
};

export default Home;
```
