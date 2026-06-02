---
id: erlang-server-sdk/scaffolds/erlang-syntax-only
sdk: erlang-server-sdk
kind: scaffold
lang: erlang
file: snippet.erl
description: |
  Parse-only validator for Erlang server SDK doc fragments.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: erlang-server
  entrypoint: snippet.erl
---

```erlang
-module(snippet).
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
