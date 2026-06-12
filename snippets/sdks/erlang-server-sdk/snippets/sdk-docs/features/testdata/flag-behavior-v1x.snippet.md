---
id: erlang-server-sdk/sdk-docs/features/testdata/flag-behavior-v1x
sdk: erlang-server-sdk
kind: reference
lang: erlang
description: Configuring test data flag behavior for Erlang SDK v1.x.
validation:
  scaffold: erlang-server-sdk/scaffolds/erlang-syntax-only-block
---

```erlang
%% This flag is true for the user key "example-user-key" and false for everyone else
{ok, Flag2} = ldclient_testdata:flag("flag-key-456def"),
UpdatedFlag2 = ldclient_flagbuilder:fallthrough_variation(false,
  ldclient_flagbuilder:variation_for_user(<<"example-user-key">>, true, Flag2)),

%% This flag returns the string variation "green" for users who have the custom
%% attribute "admin" with a value of true, and "red" for everyone else.
{ok, Flag} = ldclient_testdata:flag("flag-key-789ghi"),
UpdatedFlag = ldclient_flagbuilder:fallthrough_variation(0,
              ldclient_flagbuilder:then_return(1,
              ldclient_flagbuilder:if_match(<<"admin">>, [true],
              ldclient_flagbuilder:variations([<<"red">>, <<"green">>], Flag)))),
ldclient_testdata:update(UpdatedFlag),
```
