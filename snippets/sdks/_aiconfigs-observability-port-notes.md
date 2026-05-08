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

