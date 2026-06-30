# OpenFeature section port notes (SDK-2627)

Source: `ld-docs-private/fern/topics/sdk/openfeature/` — `node-server.mdx`,
`java.mdx`, `dotnet-server.mdx`, `php.mdx` (`index.mdx` carries no code).

Every code block is bound to a **real provider validator**: the runners
install the OpenFeature SDK + the LaunchDarkly base SDK + the LaunchDarkly
OpenFeature provider and exercise each fragment against a live LaunchDarkly
environment (node/php install at stage time; java/dotnet have the provider
packages added to the pre-baked dependency superset in their validator
Dockerfiles).

## Scaffold shape (per language)

- `*-toplevel` — resolves the import/using fragment against the real
  packages (executes / compiles; not parse-only).
- `*-init-runner` — runs the "initialize the provider" fragment's own
  provider registration, then evaluates a flag. Java uses a Main + Runner
  companion pair (the Runner reads the client back off the global
  `OpenFeatureAPI` singleton after the body registers the provider).
- `*-runner` — pre-registers a connected provider + client + context, runs
  fragments that consume them (evaluate / track / access).
- `*-context-runner` (java, dotnet) — the "construct a context" fragments
  declare their own `context`; Java and C# forbid shadowing a local the
  scaffold already bound, so these get a variant that leaves `context` to
  the fragment and evaluates with it afterward. Node and PHP reuse the
  single runner (JS block-shadowing / PHP reassignment make a separate
  variant unnecessary).

## Snippet content fixes (real bugs the validators caught)

- **java** import fragment omitted `dev.openfeature.sdk.Client`, which the
  init fragment uses.
- **java** "access the LaunchDarkly client" declared `LDClient`, but
  `Provider.getLdClient()` returns `LDClientInterface`.
- **java** organization-context fragment had a misplaced generic
  (`new HashMap(<String, Value>)` --> `new HashMap<String, Value>()`).
- **dotnet** import fragment omitted `using OpenFeature.Model;`
  (`EvaluationContext`).
- **dotnet** "evaluate a context" used the pre-v2 `GetBooleanValue`; the
  OpenFeature .NET v2 API (which the page targets) is `GetBooleanValueAsync`.
- **dotnet** "access the LaunchDarkly client" was missing its statement
  semicolon.

## CI

No matrix edit: the `node-server-sdk`, `php-server-sdk`, `java-server-sdk`,
and `dotnet-server-sdk` rows each validate the whole SDK with a server key,
so the new `sdk-docs/openfeature/` snippets are picked up automatically.
