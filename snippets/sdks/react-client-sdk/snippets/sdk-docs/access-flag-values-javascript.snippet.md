---
id: react-client-sdk/sdk-docs/access-flag-values-javascript
sdk: react-client-sdk
kind: reference
lang: javascript
description: "JavaScript in section \"Access flag values\""
validation:
  scaffold: react-client-sdk/scaffolds/react-syntax-only
---

```js
import { withLDConsumer } from 'launchdarkly-react-client-sdk';

const Home = ({ flags, ldClient /*, ...otherProps */ }) => {

  // You can call any of the methods from the JavaScript SDK
  // ldClient.identify({...})

  return flags.devTestFlag ? <div>Flag on</div> : <div>Flag off</div>;
};

export default withLDConsumer()(Home);
```
