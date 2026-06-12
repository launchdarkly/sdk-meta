---
id: erlang-server-sdk/scaffolds/erlang-syntax-only-stmts
sdk: erlang-server-sdk
kind: scaffold
lang: erlang
file: hello_erlang_server.erl
description: |
  Parse-only validator for Erlang server SDK doc fragments that are
  complete statement sequences ending with their own terminating dot
  (e.g. `LdOptions = #{...},` followed by
  `ldclient:start_instance(...).` on the storing-data pages). The
  plain `erlang-syntax-only` scaffold appends the clause-terminating
  dot itself, which would double-terminate these bodies; this variant
  lets the body's final dot close the clause.

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

%% wrappee_/0 hosts the snippet body, which supplies the terminating
%% dot itself.
wrappee_() ->
{{ body }}
```
