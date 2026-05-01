---
id: erlang-server-sdk/sdk-docs/migration-1-to-2-maps-without-kind-create-users-not-contexts-impact-of-context-kind
sdk: erlang-server-sdk
kind: reference
lang: erlang
description: "Impact of context kind in section \"Maps without \"kind\" create users, not contexts\""
---

```erlang
%% The SDK is still capable of evaluating user maps, for instance:
#{
  key => <<"example-user-key">>,
  first_name => <<"Sandy">>
}

%% If you add a kind to your map, then it will NOT be valid,
%% because `first_name` is no longer a built-in attribute for contexts.
%% Instead, create your context using ldclient_context:new
Context = ldclient_context:set(<<"firstName">>, <<"Sandy">>,
  ldclient_context:new(<<"example-user-key">>))
%% Or, use a map
Context = #{
  key => <<"example-user-key">>,
  kind => <<"user">>
  attributes => #{
    <<"firstName">> => <<"Sandy">>
  }
}
```
