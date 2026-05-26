# ai-configs and observability port notes

Validation coverage notes for the `ai-configs/`, `ai-configs-agents/`,
and `observability/` snippet groups added in PRs #424, #425, and the
agent-integration follow-up port.

This file is the analogue of `_sdk-info-port-notes.md` and
`_sdk-docs-port-notes.md`. Each entry below is a snippet (or family of
snippets) that ended up Bucket C — present in the tree, byte-checked
through the marker-hash machinery, but with no `validation:` block
because the scaffold or harness work needed to bind it cleanly is out
of scope for this PR.

## .NET ai-configs install fragment uses NuGet PowerShell cmdlets

**Severity**: low

**Snippets affected**: `dotnet-server-sdk/ai-configs/install`.

**Why unbindable**: the body is two `Install-Package` cmdlets — the
NuGet PowerShell host's package-install verbs, not the modern `dotnet
add package` CLI. The `shell-install` validator sniffs the leading
token of the body to pick a strategy (npm/pnpm/yarn/pip/go/bower/gem)
and rejects unknown leading tokens. Adding PowerShell support would
require pulling pwsh + the NuGet PowerShell cmdlets into the
shell-install image; that's a meaningful expansion of the validator's
toolchain footprint for one snippet. The sdk-info equivalent
(`install-csharp.txt`) is documented the same way in
`_sdk-info-port-notes.md`.

**Recommended action**: when the consumer refactor lands, consider
either (a) adding a parallel `install-dotnet-cli` snippet using
`dotnet add package`, which the shell-install harness could
handle by adding a `dotnet` case, or (b) deprecating the NuGet
PowerShell variant in favour of the CLI.

## ai-configs-agents Python framework full-example snippets

**Severity**: low

**Snippets affected**: all `python-server-sdk/ai-configs-agents/*`
entries — `langchain-{openai,anthropic,bedrock,gemini}`, `strands`,
`openai-agents`, and `claude`, in both `framework-install` and
`full-example` flavors.

**Why unbindable**: each `full-example` body imports a third-party
agent framework (`langchain`, `strands`, `agents` from openai-agents,
`claude_agent_sdk`) and calls a paid external API. The Python
syntax-only scaffold won't import these packages, and the runtime
scaffold can't get real API keys for OpenAI / Anthropic / Bedrock /
Gemini / Anthropic-Claude in CI. The `framework-install` snippets are
lone `pip install` lines that the `shell-install` validator could in
principle resolve, but they're intentionally fragments meant to be
appended to the base `python-server-sdk/ai-configs/install` lines and
their standalone meaning is awkward.

**Source of truth**: ported from
`static/ld/components/AiConfigs/AgentControlOnboarding/agentIntegrationSnippets.ts`
in gonfalon (see `getFullPythonAgentExample` and
`getPythonFrameworkInstallLines`). Jason flagged these as the next
slice to migrate after gonfalon#62560 (sdk-meta canonical AI Configs
adoption) lands.

**Deliberate divergences from the gonfalon source** (caught by Cursor
Bugbot on PR #444; gonfalon-only consumers of this content should
pull these fixes back when they re-sync):

- Model name lookup in LangChain / OpenAI Agents / Strands
  full-examples: gonfalon's source uses
  `config.model.get_parameter("name")`, but `name` is a top-level
  attribute on `ModelConfig` (see the canonical
  `python-server-sdk/ai-configs/implementation` snippet, where
  `ModelConfig(name=..., parameters={"temperature": 0.8})` shows
  `parameters` is only the temperature-style dict). `get_parameter("name")`
  returns `None`, making `model = ... if config.model else default`
  silently resolve to `None` whenever a variation has a model set.
  Switched to `config.model.name` to match the Claude full-example
  (which was already correct in the source) and the canonical
  `ModelConfig` shape.
- Strands `handle_agent_call_strands` was declared `async` in the
  source but its body has no `await` / `async for` (Strands `Agent`
  instances are sync-callable in the common usage). Dropped the
  `async` keyword; users who want a true async wiring can switch to
  `agent.invoke_async(...)` themselves.
- Strands `framework-install`: gonfalon's source ran
  `pip install strands openai`, but the AWS Strands Agents SDK's PyPI
  distribution is `strands-agents`, not `strands` (the `strands`
  package on PyPI is an unrelated project). The Python *import*
  namespace is `strands` — hence `from strands import Agent` in the
  full-example body — but the pip package name differs. Switched the
  install line to `pip install strands-agents openai`.
- Null-handling of `config.instructions` in the agent `system_prompt`
  / `instructions` slot: gonfalon's source only had the defensive
  `config.instructions or ""` guard on the Claude full-example;
  the four LangChain variants, OpenAI Agents, and Strands passed
  `config.instructions` raw. Aligned the other six to match Claude.
  All four underlying frameworks accept `Optional[str]` here so
  neither form crashes — `None` means "no system prompt", `""` means
  "empty system prompt" — but for a canonical copy-paste example the
  defensive form is friendlier: an unconfigured variation can’t
  surface a framework-specific complaint about `None` reaching some
  internal call. In normal use both branches are unreachable because
  the LD onboarding always seeds an instructions string.

**Open questions left for Jason / the AI Configs team**:

- `from ldai import LDAIClient, AIAgentConfig, AIAgentConfigDefault`
  in every `full-example` body — the existing canonical
  `python-server-sdk/ai-configs/import` uses
  `from ldai.client import LDAIClient`. Confirm the public
  `ldai`-top-level import path for the agent API before consumers wire
  these snippets, and align with the completion-API snippets either way.
- The Claude default model `claude-sonnet-4-6-20260217` is the value
  carried over from gonfalon's source. Confirm vs. the latest
  ai-config-quickstart docs page before publishing.
- LangChain default model ids (`openai:gpt-5`,
  `anthropic:claude-sonnet-4-6`, `bedrock/...`, `google_genai:...`)
  and the OpenAI/Strands default `gpt-5` were copied as-is; these
  fallback ids are only used when the AI Config variation doesn't set
  a model, but they should still track whatever the docs recommend.

**Recommended action**: leave Bucket C for the first port so the
canonical text lives in sdk-meta. A later slice can add an
`agent-syntax-only` Python scaffold per integration that installs the
framework's package but stubs the network call (similar to how the
existing `python-syntax-only` scaffold imports `ldclient` without
talking to LD), which would let `full-example` snippets be
import-and-parse-checked at least.

