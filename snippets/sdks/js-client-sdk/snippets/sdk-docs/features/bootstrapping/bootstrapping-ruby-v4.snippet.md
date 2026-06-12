---
id: js-client-sdk/sdk-docs/features/bootstrapping/bootstrapping-ruby-v4
sdk: js-client-sdk
kind: reference
lang: javascript
description: Bootstrapping from a Ruby template directive for JavaScript SDK v4.x.
# Not validated: the body embeds a Ruby ERB template directive
# (`<%= ... %>`) inside the JavaScript, so it is not parseable by any
# JavaScript-family validator. Wiring it up needs a mixed-host harness
# that pre-renders (or stubs out) the ERB directive before handing the
# remainder to a JS parser. The marker-hash byte-equality check in the
# docs repo is the strongest verification currently available.
---

```js
import { createClient } from '@launchdarkly/js-client-sdk';

const context = { kind: 'user', key: 'example-user-key'};

const client = createClient(
  'example-client-side-id', 
  context
);

client.start({
  // Load values from a Ruby template directive
  bootstrap: <%= client.all_flags_state(user, {client_side_only: true}).to_json %>
});
```
