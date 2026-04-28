---
id: erlang-server-sdk/getting-started/server-erl
sdk: erlang-server-sdk
kind: hello-world
lang: erlang
file: src/hello_erlang_server.erl
description: gen_server module that wraps the LaunchDarkly Erlang client.
inputs:
  apiKey:
    type: sdk-key
    description: SDK key baked into the rendered source. Validation is not yet wired — see comment.
ld-application:
  slot: server-erl
# Validator pending. The Erlang Get Started flow is fundamentally
# interactive (rebar3 shell + manual gen_server:call). To validate end-
# to-end we'd need to write a wrapper application script that calls
# hello_erlang_server:get/3 and prints an EXAM-HELLO conformant line.
# That's a snippet rewrite and is left as fix-on-red.
---

First create a new file named `src/hello_erlang_server.erl`. Then, in
`src/hello_erlang_server.erl` create a new `LDClient` with your *environment-specific* SDK key:

```erlang
-module(hello_erlang_server).
-behaviour(gen_server).

-export([init/1, handle_call/3, handle_cast/2,
         handle_info/2, terminate/2, code_change/3]).

-export([start_link/0]).
-export([get/3]).

% public functions

start_link() ->
  gen_server:start_link({local, hello_erlang_server}, ?MODULE, [], []).

get(Key, Fallback, ContextKey) -> gen_server:call(?MODULE, {get, Key, Fallback, ContextKey}).

% gen_server callbacks

init(_Args) ->
  ldclient:start_instance("{{ apiKey }}", #{
        http_options => #{
            tls_options => ldclient_config:tls_basic_options()
        }
    }),
  {ok, []}.

handle_call({get, Key, Fallback, ContextKey}, _From, State) ->
  Flag = ldclient:variation(Key, ldclient_context:new(ContextKey), Fallback),
  {reply, Flag, State}.

handle_cast(_Request, State) ->
  {noreply, State}.

handle_info(_Info, State) ->
  {noreply, State}.

terminate(_Reason, _State) ->
  ok.

code_change(_OldVsn, State, _Extra) ->
  {ok, State}.
```
