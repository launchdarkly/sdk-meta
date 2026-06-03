---
id: erlang-server-sdk/sdk-docs/initialize-the-client-erlang
sdk: erlang-server-sdk
kind: reference
lang: erlang
description: "Erlang in section \"Initialize the client\""
# TODO(snippet-bug): body has two `ldclient:start_instance(...)`
# expression statements without a `,` separator. Erlang requires
# `,` between statements inside a function body — without one, the
# parser bails at "syntax error before: ldclient". Fix in the
# snippet-bugs PR: add the `,` (or render as two separate snippets).
---

```erlang
% This starts an instance with the default options
ldclient:start_instance("YOUR_SDK_KEY")

% You can also start a named instance
ldclient:start_instance("YOUR_SDK_KEY", your_instance)
```
