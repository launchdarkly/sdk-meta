---
id: erlang-server-sdk/scaffolds/erlang-syntax-only
sdk: erlang-server-sdk
kind: scaffold
lang: erlang
file: hello_erlang_server.erl
description: |
  Parse-only validator for Erlang server SDK doc fragments.

  Stages as `hello_erlang_server.erl` to match the filename the
  erlang-server harness hard-codes when copying snippets into the
  pre-baked rebar3 project (Erlang requires `-module(Name)` to
  match the source file's basename, so scaffold and harness must
  agree on the name).

  The harness dispatches on `SNIPPET_CHECK=parse` to skip the
  runtime `erl` invocation that the init-runner variant uses, and
  just confirms the staged file compiles cleanly.
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

%% wrappee/0 hosts the snippet body. Erlang requires function names
%% to start with a lowercase letter (atoms), so the conventional
%% leading-underscore name from other scaffolds doesn't translate
%% here — `wrappee_` is the closest lexical match.
wrappee_() ->
{{ body }}.
```
