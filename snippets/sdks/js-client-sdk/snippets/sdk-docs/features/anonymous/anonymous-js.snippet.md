---
id: js-client-sdk/sdk-docs/features/anonymous/anonymous-js
sdk: js-client-sdk
kind: reference
lang: javascript
description: Anonymous context example for JavaScript, SDK v3.0.
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only
---

```javascript
const anonymousUserContext = {
  kind: 'user',
  anonymous: true
};

// A multi-context can contain both anonymous and non-anonymous contexts.
// Here, the organization is not anonymous.
const multiContext = {
  kind: 'multi',
  user: anonymousUserContext,
  org: {
    key: 'example-organization-key',
    name: 'Acme, Inc.'
  }
}
```
