# Port notes: /sdk/features/migrations

Source: `ld-docs-private` `fern/topics/sdk/features/migrations.mdx`.
46 code blocks extracted into `sdk-docs/features/migrations/`
snippets across 11 SDKs (8 server-side: .NET, Go, Java, Node.js, PHP,
Python, Ruby, Rust; 3 edge: Akamai, Cloudflare, Vercel). All 46 are
bound to validators.

## Content corrections

Fixes applied where the published sample could not work as written.
Everything else is verbatim from the MDX (modulo indentation flush to
column 0 and leading/trailing blank-line trim inside fences).

- **.NET read/write** (`dotnet-server-sdk/.../read-write`) and
  **.NET MigrationVariation** (`dotnet-server-sdk/.../migration-variation`):
  `LDContext context = Context.Builder(...)` -- the .NET SDK has no
  `LDContext` type (that's the Java/Android name); the context struct
  is `LaunchDarkly.Sdk.Context`. Changed the declared type to
  `Context`.
- **.NET stage switch** (`dotnet-server-sdk/.../stage-switch`): the
  final switch section (the stacked case labels plus `default:`)
  contained only a comment. A C# switch section must contain at least
  one statement and may not fall out, so the block did not compile.
  Added `break;` after the `// throw an error` comment.
- **Node.js read/write** (`node-server-sdk/.../read-write`): the
  import list contained `LD MigrationState` -- a syntax error (stray
  space) and a type that does not exist. The package exports
  `LDMigrationStage`, which is what the body's
  `let defaultStage: LDMigrationStage = ...` declaration uses.
- **Node.js MigrationVariation** (`node-server-sdk/.../migration-variation`):
  the import line ended `from '@launchdarkly/node-server-sdk');` --
  a stray `)` that is a syntax error. Also, `migrationVariation` was
  called with `false` as the default stage; the parameter is an
  `LDMigrationStage`, so the call now passes `LDMigrationStage.Off`
  (added to the import list).
- **Akamai / Cloudflare / Vercel MigrationVariation**
  (`akamai-server-edgekv-sdk`, `cloudflare-server-sdk`,
  `vercel-server-sdk` `.../migration-variation`): same wrong default
  as Node -- `false` passed where an `LDMigrationStage` is expected.
  Changed to `LDMigrationStage.Off` and added it to each import list.
- **Node.js / Akamai / Cloudflare / Vercel stage switch**
  (`.../stage-switch`, 4 snippets): every case clause read
  `case LDMigrationStage.Off: { },` -- a comma after a block
  statement, which is a syntax error in JavaScript/TypeScript.
  Removed the trailing commas.
- **Python stage switch** (`python-server-sdk/.../stage-switch`): the
  `if/elif` chain had no statements in any branch and the `else:`
  branch held only a comment -- a SyntaxError on any Python. Added
  `pass` to each branch, keeping the page's two-space indentation.

## Validation routing added in this port

- `dotnet-server-sdk/scaffolds/csharp-syntax-only-typed` -- variant of
  `csharp-syntax-only` whose `client` stub is a real `LdClient`
  instead of `dynamic`. C# rejects tuple deconstruction of a dynamic
  result (CS8133), so
  `var (stage, tracker) = client.MigrationVariation(...)` cannot
  compile through the dynamic stub. The same scaffold id (and shape)
  exists on the in-flight monitoring port branch for the CS1977
  lambda-through-dynamic case; this copy adds the
  `LaunchDarkly.Sdk.Server.Migrations` using that the migration body
  needs, and the descriptions document both blockers.
- Stub-surface extensions:
  - `csharp-syntax-only`: `migration`, `payload`, `tracker` (dynamic)
    and `stage` (`MigrationStage`) fields for the ambient names the
    migration fragments reference.
  - `java-syntax-only`: `migration`
    (`Migration<String, String, String, String>`) and
    `migrationVariation` (`MigrationVariation`) stub fields.
  - `rust-syntax-only`: `Stage` and `MigrationOpTracker` added to the
    SDK use list, `Mutex` alongside `Arc`; a `_StubMigrator` type with
    concrete-typed async `read`/`write` methods (the real `Migrator`
    is generic over its closure types, which a stub binding cannot
    name, and concrete parameter types let the body's `.into()` calls
    infer); `migrator`, `stage`, and `tracker` bindings in
    `_wrappee`.
- `validators/languages/rust/Dockerfile`: base image bumped
  `rust:1.85` --> `rust:1.94`. `cargo add launchdarkly-server-sdk`
  resolves the newest SDK the toolchain's MSRV allows, and the 3.x
  line requires 1.94. Identical change is on the in-flight flush port
  branch.

## Known non-binds

None -- every block on this page is bound. (The page covers only
server-side and edge SDKs; the iOS Objective-C blocker from earlier
ports does not arise here.)
