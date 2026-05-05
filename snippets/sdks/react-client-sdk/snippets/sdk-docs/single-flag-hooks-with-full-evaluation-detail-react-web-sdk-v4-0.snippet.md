---
id: react-client-sdk/sdk-docs/single-flag-hooks-with-full-evaluation-detail-react-web-sdk-v4-0
sdk: react-client-sdk
kind: reference
lang: javascript
description: "React Web SDK v4.0 in section \"Single-flag hooks with full evaluation detail\""
---

```js
import { useBoolVariationDetail } from '@launchdarkly/react-sdk';

const { value, variationIndex, reason } = useBoolVariationDetail('example-flag-key', false);
```
