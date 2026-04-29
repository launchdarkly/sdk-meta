FROM erlang:26

# rebar3 ships with the official erlang image.
#
# Pre-bake a hello_erlang OTP application matching the structure
# gonfalon's Get Started flow walks the user through:
#   - rebar.config pulls launchdarkly_server_sdk from Hex (alias ldclient).
#   - hello_erlang.app.src lists ldclient as an application.
#   - hello_erlang_sup.erl supervises hello_erlang_server.
#
# Per-validate cycles only swap hello_erlang_server.erl, run a small
# rebar3 compile (incremental against the warmed _build), and execute a
# rebar3 eval that calls into the gen_server and prints the EXAM-HELLO
# canonical line. The snippet itself is a gen_server with no main entry
# point — gonfalon's flow expects the user to drop into rebar3 shell
# and call manually. We synthesize the equivalent at validate time.

ARG LD_ERLANG_VERSION=3.9.0

WORKDIR /opt/hello_erlang

RUN rebar3 new app hello_erlang \
    && mv hello_erlang/* hello_erlang/.??* . 2>/dev/null || true \
    && rmdir hello_erlang 2>/dev/null || true

# Pin the SDK + add it to the app-src + supervisor child spec. We bake
# the same content the snippet's manifest-fragments would otherwise
# describe.
RUN printf '{erl_opts, [debug_info]}.\n\
{deps, [\n\
  {ldclient, "%s", {pkg, launchdarkly_server_sdk}}\n\
]}.\n\
{shell, [\n\
  {apps, [hello_erlang]}\n\
]}.\n' "${LD_ERLANG_VERSION}" > rebar.config

RUN printf '{application, hello_erlang,\n\
 [{description, "Hello LaunchDarkly"},\n\
  {vsn, "0.1.0"},\n\
  {registered, []},\n\
  {mod, {hello_erlang_app, []}},\n\
  {applications,\n\
   [kernel,\n\
    stdlib,\n\
    ldclient\n\
   ]},\n\
  {env,[]},\n\
  {modules, []}\n\
 ]}.\n' > src/hello_erlang.app.src

RUN cat > src/hello_erlang_sup.erl <<'EOF'
-module(hello_erlang_sup).
-behaviour(supervisor).
-export([start_link/0]).
-export([init/1]).
-define(SERVER, ?MODULE).

start_link() ->
    supervisor:start_link({local, ?SERVER}, ?MODULE, []).

init([]) ->
    SupFlags = #{strategy => one_for_all, intensity => 0, period => 1},
    ChildSpecs = [{console,
            {hello_erlang_server, start_link, []},
            permanent, 5000, worker, [hello_erlang_server]}],
    {ok, {SupFlags, ChildSpecs}}.
EOF

# Placeholder server module — replaced per-validate.
RUN cat > src/hello_erlang_server.erl <<'EOF'
-module(hello_erlang_server).
-behaviour(gen_server).
-export([start_link/0, init/1, handle_call/3, handle_cast/2,
         handle_info/2, terminate/2, code_change/3]).
-export([get/3]).
start_link() -> gen_server:start_link({local, ?MODULE}, ?MODULE, [], []).
get(_, F, _) -> F.
init(_) -> {ok, []}.
handle_call(_, _, S) -> {reply, ok, S}.
handle_cast(_, S) -> {noreply, S}.
handle_info(_, S) -> {noreply, S}.
terminate(_, _) -> ok.
code_change(_, S, _) -> {ok, S}.
EOF

RUN rebar3 compile

WORKDIR /work

COPY shared /harness-shared
COPY languages/erlang-server/harness /harness

ENTRYPOINT ["/harness/run.sh"]
