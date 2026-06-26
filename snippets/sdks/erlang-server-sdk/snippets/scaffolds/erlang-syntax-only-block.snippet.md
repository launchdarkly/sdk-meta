---
id: erlang-server-sdk/scaffolds/erlang-syntax-only-block
sdk: erlang-server-sdk
kind: scaffold
lang: erlang
file: hello_erlang_server.erl
description: |
  Statement-sequence sibling of `erlang-syntax-only`. That scaffold
  closes the spliced body with a bare `.`, which only works when the
  body is a complete expression with no trailing separator. Doc
  fragments that are excerpts of a larger function body end each
  expression with `,` (including the last line), so appending `.`
  directly produces the invalid token sequence `,.`. This variant
  closes the sequence with a final `{ok, SdkKey}` expression instead,
  so trailing-comma bodies compile unchanged.

  `SdkKey` is bound at the top of the function because config-style
  fragments pass an ambient `SdkKey` the docs assume an enclosing
  function provided; Erlang treats unbound variables as compile
  errors, unlike unresolved remote calls. Referencing it again in the
  closing expression keeps the binding warning-free for bodies that
  never use it.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, a comma-terminated expression sequence.
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

%% wrappee_/0 hosts the snippet body. The body's expressions each end
%% with `,`, so the closing expression below both terminates the
%% sequence and consumes the SdkKey binding.
wrappee_() ->
    SdkKey = <<"example-sdk-key">>,
{{ body }}
    {ok, SdkKey}.
```
