# Port notes: /sdk/features/flags-from-files

Source: `ld-docs-private` `fern/topics/sdk/features/flags-from-files.mdx`.
19 code blocks extracted into `sdk-docs/features/filedata/` snippets
across 10 server-side SDKs plus six shared page-level examples (three
curl commands, two JSON flag-data files, one YAML flag-data file). All
13 per-SDK snippets are bound to validators; the six `_shared` examples
are documented non-binds (below).

## Content corrections

Fixes applied where the published sample could not work as written.
Everything else is verbatim from the MDX (modulo indentation flush to
column 0).

- **.NET (server) configuration** (`dotnet-server-sdk/.../flags-from-files`):
  three errors — the builder chain ended in `.Build()` with no
  semicolon (a syntax error); `Components.NoEvents()` called a static
  property as a method (it is `Components.NoEvents`); and
  `new LDClient(config)` used the Java class name — the .NET server
  SDK's client class is `LdClient`.
- **Node.js v8 examples** (`node-server-sdk/.../flags-from-files-js-v8`,
  `-ts-v8`): both import lines read
  `from '@launchdarkly/node-server-sdk');` — a stray `)` left over
  from converting the v7 `require(...)` form, and a syntax error.
  Fixed to `from '@launchdarkly/node-server-sdk';`.

## Validation routing added in this port

- `erlang-server-sdk/scaffolds/erlang-syntax-only` — the clause
  terminator after `{{ body }}` moved onto its own line. The Erlang
  fragment on this page ends with `%%`-comment lines; a dot appended
  directly after the body is swallowed by the trailing comment and the
  clause never terminates. A stand-alone dot terminates the clause for
  comment-ending and expression-ending bodies alike. All eleven bound
  erlang wrappees revalidated.
- `haskell-server-sdk/scaffolds/haskell-syntax-only` — added qualified
  aliases `import qualified LaunchDarkly.Server as LD` and
  `import qualified LaunchDarkly.Server.Integrations.FileData as FileData`;
  the page's fragment references `LD.` / `FileData.` without showing
  its own import lines. All five bound wrappees revalidated.
- `cpp-server-sdk/scaffolds/cpp-syntax-only-v2-c` — the scaffold now
  pre-includes `<launchdarkly/integrations/file_data.h>` at file scope.
  The fragment's own `#include` sits inside the `_wrappee()` body where
  the header's `static inline` definitions must not land; with the
  guard already defined, the in-body include expands to nothing.
- `validators/languages/cpp-server-v2-c` — new stub
  `integrations/file_data.h` (declares `LDFileDataInit`, mirroring the
  real v2 header's `struct LDDataSource *(int, const char **)`
  signature) staged into the image by the Dockerfile; `api.h` gained
  `struct LDDataSource`, `LDConfigSetDataSource`, and
  `LDConfigSetSendEvents`. The `cpp-server-v2-cpp` validator shares
  `api.h`, so its one wrappee was revalidated too.

## Known non-binds

Six `_shared` page-level examples (`sdk-docs/features/filedata/...`),
all blocked the same way as the identify port's shared JSON
multi-context example: `snippets validate` requires `--sdk` and
filters on the frontmatter `sdk:` field, which `_shared` snippets do
not carry — no CI row can currently select them. Per-format blockers
on top of that:

- `flag-data-json`, `flag-values-json`, `flag-values-yaml` — no
  `json` / `yaml` validator runtime exists yet (a
  `python3 -m json.tool` / PyYAML-style parse harness would do).
- `flag-data-curl` (+`-federal`, `-eu`) — the `shell-install` harness
  only dispatches on package-manager leading tokens (`npm`, `pip`,
  `go`, ...) and asserts install side effects; `curl` is a hard error
  there. Running the commands for real would also need live
  credentials per instance (commercial, federal, EU) — CI provisions
  only the commercial sandbox key.
