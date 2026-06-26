# Port notes: /sdk/features/agentcontrol-config

Source: `ld-docs-private` `fern/topics/sdk/features/agentcontrol-config.mdx`.
7 code blocks extracted into `sdk-docs/features/agentcontrol/` snippets
across 5 SDKs (.NET, Go, Node.js, Python, Ruby — the page covers the AI
SDKs only). All 7 are bound to validators.

## Content corrections

Fixes applied where the published sample could not work as written.
Everything else is verbatim from the MDX. API claims were verified
against the local SDK checkouts (`dotnet-core/pkgs/sdk/server-ai`,
`go-server-sdk/ldai`, `js-core/packages/sdk/server-ai`,
`python-server-sdk-ai`, `ruby-server-sdk-ai`).

- **.NET customize example** (`dotnet-server-sdk/.../customize-config`):
  the published sample used the pre-0.10.0 API. Since
  LaunchDarkly.ServerSdk.Ai 0.10.0, `LdAiConfig` is the abstract
  SDK-output type — the fallback builder moved to
  `LdAiCompletionConfigDefault.New()`, and `Config(...)` returns the
  config itself (`LdAiCompletionConfig`) rather than a tracker with a
  `.Config` property (trackers now come from `config.CreateTracker()`).
  Rewritten to build the fallback via `LdAiCompletionConfigDefault`,
  name the result `config`, and test `config.Enabled`. Also added the
  missing semicolon after `.Build()`.
- **Go customize example** (`go-server-sdk/.../customize-config`): the
  fallback was built with an unqualified `NewConfig()`, which does not
  resolve in user code — the builder lives in the `ldai` package
  (`ldai.NewConfig()`), matching the `ldai.Disabled()` reference in the
  surrounding prose.
- **Node.js agent-mode example** (`node-server-sdk/.../agent-config`):
  the published block was fenced as Python and glued a Python
  `LDAIAgentConfig(...)` construction onto the TypeScript
  `aiClient.agentConfig(...)` call, which referenced an undefined
  `fallbackConfig`. Replaced the Python half with a
  `const fallbackConfig = { enabled: false };` declaration paralleling
  the completion-mode block; the TypeScript half is kept as published.
  The MDX fence language changes from `python` to `ts` with this fix.
- **Python completion-mode example** (`python-server-sdk/.../completion-config`):
  `completion_config()` returns a single `AICompletionConfig` (the
  tracker rides on `config.tracker`), not a `(config, tracker)` tuple —
  the tuple unpack raises TypeError on the current SDK. Also the
  `if config.enabled:` / `else:` branches contained only comments,
  which is a SyntaxError on any Python; added `pass` to each branch
  (same approach as the evaluating port's python snippet).
- **Ruby customize example** (`ruby-server-sdk/.../customize-config`):
  `LaunchDarkly::AI::AIConfig` does not exist — the Ruby AI SDK's
  namespace is `LaunchDarkly::Server::AI` (the page's own API link
  points at `LaunchDarkly/Server/AI/Client`). Qualified the fallback as
  `LaunchDarkly::Server::AI::AIConfig.new(enabled: false)`.

Verbatim (no changes): `node-server-sdk/.../completion-config`
(`LDAIConfig` is the base interface `completionConfig`'s
`LDAICompletionConfig` return type extends, so the annotation is
valid) and `python-server-sdk/.../agent-config` (`AIAgentConfigDefault`
is importable from `ldai` and `agent_config` takes
key/context/default/variables positionally).

## Validation routing

No new scaffolds, validators, or CI rows were needed. All five SDKs
already have syntax-only scaffolds with the ambient names these bodies
use (`aiClient`, `aiclient`, `ai_client`, `context`), introduced by the
ai-configs port:

- `dotnet-server-sdk/scaffolds/csharp-syntax-only` — real compile; the
  validator's requirements already stage `LaunchDarkly.ServerSdk.Ai`,
  so the .NET body's API usage is checked against the released package.
- `go-server-sdk/scaffolds/go-syntax-only`,
  `node-server-sdk/scaffolds/node-syntax-only`,
  `python-server-sdk/scaffolds/python-syntax-only`,
  `ruby-server-sdk/scaffolds/ruby-syntax-only` — parse-only.

The existing CI matrix rows for these five SDKs carry no `--group`
filter, so the new `sdk-docs/features/agentcontrol/` snippets are
picked up without workflow changes.

## Known non-binds

None — every block on the page is bound.
