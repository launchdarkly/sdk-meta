# ai-configs and observability port notes

Validation coverage notes for the `ai-configs/` and `observability/`
snippet groups added in PRs #424 and #425.

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

## Frontend observability `import` fragments need the o11y packages
on the validator's dep list

**Severity**: medium

**Snippets affected**:
`js-client-sdk/observability/import`,
`react-client-sdk/observability/import`,
`react-native-client-sdk/observability/import`,
`vue-client-sdk/observability/import`.

**Why unbindable**: each body has top-level `import Observability,
{ LDObserve } from '@launchdarkly/observability'` (and, for
js/react/vue, a parallel `@launchdarkly/session-replay` import). The
existing per-SDK syntax-only scaffolds either (a) wrap the wrappee
inside a function body, where ES `import` statements are not legal,
or (b) stage the body at module scope but in a project whose
package.json declares only the core SDK — the bundler (tsdown for
js-client, Vite for react/vue, jest+RN preset for react-native) fails
to resolve the o11y package paths. The sdk-info `import` snippets
don't hit this because they import only from the core SDK package
that the validator already ships.

The python and node-server `observability/import` snippets DO bind
cleanly to their syntax-only scaffolds (Python's `ast.parse` and
Node's `node --check` are both purely syntactic — they don't try to
resolve the import paths).

**Recommended action**: extend each frontend validator's Dockerfile
to install the relevant `@launchdarkly/observability*` package (and
session-replay where applicable), and add an `IMPORT_LIFT_TARGET` /
`BODY_BEGIN` / `BODY_END` marker pair to each affected scaffold so
the harness can hoist the body's `import` lines to module scope (the
react-client harness already implements this pattern and is a good
reference). One follow-up PR per validator keeps the diffs small.

## Observability `initialize` fragments need init-runner scaffold
extension + o11y deps on the validator

**Severity**: medium

**Snippets affected**:
`js-client-sdk/observability/initialize`,
`node-server-sdk/observability/initialize`,
`python-server-sdk/observability/initialize`,
`react-client-sdk/observability/initialize`,
`react-native-client-sdk/observability/initialize`,
`vue-client-sdk/observability/initialize`.

**Why unbindable**: each body assumes the symbols introduced by the
companion `import.snippet.md` (`init`, `Observability`,
`SessionReplay`, `LDPlugin`, `withLDProvider`, `createApp`,
`AutoEnvAttributes`, etc.) are already in scope, and shows only the
plugin-options portion of the SDK initialization call. To run
end-to-end through an init-runner scaffold, every body would need
the matching `import` block prepended before the body is staged
into the runner — a level of pre-processing the existing init-runner
scaffolds don't do today.

The bodies also use a literal `'SDK_KEY'` placeholder (rather than
the per-SDK `'YOUR_SDK_KEY'` / `'YOUR_CLIENT_SIDE_ID'` /
`'YOUR_MOBILE_KEY'` strings the existing init-runner harnesses
substitute via `validation.placeholders`). The placeholder name is
fine in isolation, but the per-SDK init-runners were calibrated to
the sdk-info init bodies' specific placeholder strings.

Finally, the runtime image needs the observability + session-replay
packages installed alongside the core SDK. The python and node-server
init-runners pull dependencies through `validation.requirements` /
`requirements.txt`, so adding `launchdarkly-observability` and
`@launchdarkly/observability-node` is a one-line change there. The
browser/RN init-runners pre-bake the SDK in the Dockerfile, so each
of those Dockerfiles needs an extra dep declaration plus an image
rebuild.

**Recommended action**: introduce a companion init-runner variant
per SDK (`init-runner-with-o11y` or similar) that:

  1. Prepends the SDK + observability imports to the spliced body.
  2. Maps the `'SDK_KEY'` placeholder to the correct env var
     (LAUNCHDARKLY_SDK_KEY for server, LAUNCHDARKLY_CLIENT_SIDE_ID
     for browser, LAUNCHDARKLY_MOBILE_KEY for react-native).
  3. Wraps the body's call so the scaffold can wait for the
     resulting client's `waitForInitialization()` (or equivalent)
     to resolve and print the EXAM-HELLO sentinel.

Each variant is ~20-30 lines of YAML + harness work; collectively
manageable as a follow-up PR once a single reference SDK has been
worked through end-to-end.
