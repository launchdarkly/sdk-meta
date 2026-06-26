# Port notes: /sdk/features/ai-metrics

Source: `ld-docs-private` `fern/topics/sdk/features/ai-metrics.mdx`.
55 code blocks extracted into `sdk-docs/features/aimetrics/` snippets
across 5 SDKs. All 55 are bound to validators.

The page covers the AI SDKs only (.NET AI, Go AI, Node.js server AI,
Python AI, Ruby AI). Following the context-config port's precedent for
AI SDK blocks, each snippet is hosted under the corresponding server
SDK directory (`dotnet-server-sdk`, `go-server-sdk`,
`node-server-sdk`, `python-server-sdk`, `ruby-server-sdk`). The
context-config port disambiguated AI blocks with `-ai` slug suffixes
because they shared a folder with non-AI blocks; here the entire
`features/aimetrics/` folder is AI-SDK content, so the slugs carry no
suffix. The page's intro example (delayed feedback via
`ldClient.track`) is Node AI SDK code and lives under
`node-server-sdk`.

## Content corrections

Fixes applied where the published sample could not work as written.
API claims verified against the local SDK repos (`js-core`
`packages/sdk/server-ai` at the page's pre-1.0 target API,
`python-server-sdk-ai`, `ruby-server-sdk-ai`, `go-server-sdk/ldai`,
`dotnet-core/pkgs/sdk/server-ai`). Everything else is verbatim from
the MDX.

- **Node agent mode** (`node-server-sdk/.../agent-instructions`):
  `agent.tracker.track_success()` is the Python AI SDK's snake_case
  name pasted into a JavaScript block; the Node tracker method has
  always been `trackSuccess()`.
- **Node Vercel generateText** (`node-server-sdk/.../vercel-generate-text`):
  `trackVercelAISDKGenerateTextMetrics` takes a function returning the
  operation's promise (`func: () => Promise<TRes>`) at every version
  that shipped it; the published block passed the `generateText(...)`
  promise directly. Wrapped the call in `() =>`.
- **Node Vercel streamText** (`node-server-sdk/.../vercel-stream-text`):
  same shape — `trackVercelAISDKStreamTextMetrics(func: () => TRes)`
  takes a function; wrapped `streamText(...)` in `() =>`.
- **Python completion/agent mode** (`python-server-sdk/.../openai-completion`,
  `bedrock-completion`, `agent-instructions`): each block ended with an
  `else:` whose body was only a comment — an IndentationError on every
  Python version (comments are not statements). Added `pass` under the
  comment.
- **Python agent mode** (`python-server-sdk/.../agent-instructions`):
  `var result = ...` is not Python syntax; dropped the stray `var`.
- **Python feedback** (`python-server-sdk/.../track-feedback`):
  `track_feedback` has always taken a dict
  (`feedback: Dict[str, FeedbackKind]`, read via `feedback["kind"]`);
  the published call passed the bare enum member. Changed to
  `track_feedback({"kind": FeedbackKind.Positive})`.
- **Ruby OpenAI / Bedrock wrappers** (`ruby-server-sdk/.../openai-completion`,
  `bedrock-completion`): `track_openai_metrics(&block)` and
  `track_bedrock_converse_metrics(&block)` take a block and no
  positional arguments; passing the provider result as an argument
  raises ArgumentError. Rewritten to `do ... end` block form.
- **Ruby token usage** (`ruby-server-sdk/.../track-tokens`): the class
  is `LaunchDarkly::Server::AI::TokenUsage` (the gem nests everything
  under `Server`); the published block used the nonexistent
  `LaunchDarkly::AI::TokenUsage`.

## Validation routing added in this port

No new scaffolds or validators. All 55 snippets bind to the five
existing syntax-only scaffolds (`csharp-syntax-only`,
`go-syntax-only`, `node-syntax-only`, `python-syntax-only`,
`ruby-syntax-only`). Stub-surface extensions:

- `dotnet-server-sdk/scaffolds/csharp-syntax-only`: added
  `using System.Threading.Tasks;` (bodies call `Task.Run`) and
  `using LaunchDarkly.Sdk.Server.Ai.Tracking;` (bodies construct the
  real `Response` / `Usage` / `Metrics` records and pass
  `Feedback.Positive`), plus dynamic stub fields `tracker`,
  `response`, and `baseClient` for the ambient names the fragments
  assume.
- `node-server-sdk/scaffolds/node-syntax-only`: the TS-erasure pass
  for `as Type` assertions consumed `(` / `)` as part of the type
  text, so parenthesized casts like
  `(aiConfig.model?.parameters?.temperature as number) ?? 0.5` lost
  their closing paren and failed `node --check`. Removed the paren
  characters from the type-match class; existing bound snippets only
  use `import * as ...` (never matched) so behavior is unchanged for
  them.

## Known non-binds

None — every block on this page is bound. (This page has no iOS
Objective-C or other unscaffolded-language blocks.)
