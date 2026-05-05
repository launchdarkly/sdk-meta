---
id: react-client-sdk/sdk-docs/single-flag-hooks-react-web-sdk-v4-0
sdk: react-client-sdk
kind: reference
lang: javascript
description: "React Web SDK v4.0 in section \"Single-flag hooks\""
---

```js
import {
  useBoolVariation,
  useStringVariation,
  useNumberVariation,
  useJsonVariation,
} from '@launchdarkly/react-sdk';

const showNewFeature = useBoolVariation('show-new-feature', false);
const theme          = useStringVariation('ui-theme', 'light');
const maxItems       = useNumberVariation('max-items', 10);
const config         = useJsonVariation('my-config', {});
```
