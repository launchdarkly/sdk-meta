---
id: react-client-sdk/sdk-docs/multi-environment-support-root-tsx
sdk: react-client-sdk
kind: reference
lang: javascript
description: "JavaScript: Root.tsx in section \"Multi-environment support\""
---

```js
function Root() {
  return (
    <ProdLDProvider>
      <StagingLDProvider>
        <App />
      </StagingLDProvider>
    </ProdLDProvider>
  );
}
```
