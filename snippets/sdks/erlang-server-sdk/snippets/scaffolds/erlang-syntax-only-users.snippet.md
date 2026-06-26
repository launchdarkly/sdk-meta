---
id: erlang-server-sdk/scaffolds/erlang-syntax-only-users
sdk: erlang-server-sdk
kind: scaffold
lang: erlang
file: hello_erlang_server.erl
description: |
  Parse-only validator for Erlang server SDK doc fragments whose
  bodies reference ambient `User` / `PreviousUser` variables the docs
  assume earlier snippets bound (the v1.x aliasing fragment, for
  example). erlc rejects unbound variables, so the plain
  `erlang-syntax-only` scaffold cannot host these bodies; this
  variant pre-binds both names immediately before the body.

  Same staging contract as `erlang-syntax-only`: the file is staged
  as `hello_erlang_server.erl` to match the module name the
  erlang-server harness hard-codes, and `SNIPPET_CHECK=parse` makes
  the harness stop after a clean `rebar3 compile`.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: erlang-server
  entrypoint: hello_erlang_server.erl
  env:
    SNIPPET_CHECK: parse
---

```erlang
-module(hello_erlang_server).
-export([main/0]).

main() ->
    io:format("feature flag evaluates to true~n").

%% wrappee_/0 hosts the snippet body. The doc fragments assume the
%% reader's earlier snippets already bound `User` / `PreviousUser`;
%% bind both here so the body compiles.
wrappee_() ->
    User = #{key => <<"user-key">>},
    PreviousUser = #{key => <<"previous-user-key">>},
{{ body }}.
```
