---
id: react-client-sdk/observability/error-boundary
sdk: react-client-sdk
kind: reference
lang: javascript
file: react-client-sdk/observability/error-boundary.txt
description: Forward errors caught by a React error boundary to LaunchDarkly observability.
validation:
  scaffold: react-client-sdk/scaffolds/react-syntax-only
---

```javascript
import React from 'react';
import { LDObserve } from '@launchdarkly/observability';

class ErrorBoundary extends React.Component {
  // ...

  componentDidCatch(error, errorInfo) {
    // Forward error to LaunchDarkly
    LDObserve.recordError(
      error,
      'React Error Boundary',
      { componentStack: errorInfo.componentStack },
    );
  }

  // ...
}
```
