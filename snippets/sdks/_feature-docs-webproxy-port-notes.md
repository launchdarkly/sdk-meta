# Port notes: /sdk/features/web-proxy

Source: `ld-docs-private` `fern/topics/sdk/features/web-proxy.mdx`.
32 code blocks extracted into `sdk-docs/features/webproxy/` snippets
across 10 SDKs. 22 are bound to validators; the 10 remaining are bare
shell env-var fragments with no fitting harness (see "Known
non-binds").

## Content corrections

Fixes applied where the published sample could not work as written.
Everything else is verbatim from the MDX (modulo indentation flush to
column 0).

- **.NET (client) proxy examples** (`dotnet-client-sdk/.../web-proxy`,
  `web-proxy-auth`): the published samples call
  `.HttpMessageHandler(handler)` directly on the configuration
  builder, but no `ConfigurationBuilder.HttpMessageHandler` method
  exists in the client SDK versions that take the
  `AutoEnvAttributes` parameter shown on the same line. A custom
  handler is configured through the HTTP configuration sub-builder;
  rewritten to `.Http(Components.HttpConfiguration().MessageHandler(handler))`.
- **.NET (server) proxy with authentication**
  (`dotnet-server-sdk/.../web-proxy-auth`): the line
  `var proxyUri = new Uri("http://my-proxy-host:8080")` was missing
  its terminating semicolon — a syntax error. Added it.
- **Node.js v7 TypeScript example**
  (`node-server-sdk/.../web-proxy-ts-v7`): the published block used a
  default import (`import LDOptions from 'launchdarkly-node-server-sdk'`),
  but `LDOptions` is a named interface export in the v7 typings — a
  default import binds a value, which cannot be used in the
  `const options: LDOptions` type annotation. Changed to the named
  import, matching the authenticated v7 TypeScript block on the same
  page.

## Validation routing added in this port

No new scaffolds. Stub-surface and validator extensions:

- `validators/languages/cpp-client-v2-c/api.h` — added
  `LDConfigSetProxyURI` (returns `LDBoolean`, matching the real v2
  `launchdarkly/c-client-sdk` header) so the v2 C client proxy
  fragments type-check.
- `dotnet-client-sdk/scaffolds/csharp-client-syntax-only` and
  `dotnet-server-sdk/scaffolds/csharp-syntax-only` — added
  `using System.Net;` to the scaffold preamble. The proxy-auth
  fragments construct `NetworkCredential` unqualified (while fully
  qualifying `System.Net.WebProxy` / `System.Net.CredentialCache`),
  assuming an ambient using directive.
- `validators/languages/rust/Dockerfile` — base image `rust:1.85` -->
  `rust:1.94`. Cargo's MSRV-aware resolver picks the newest
  `launchdarkly-server-sdk` the toolchain supports; the 3.x line
  (which ships the transport layer the proxy fragment configures)
  requires 1.94.
- `validators/languages/rust/harness/run.sh` — also `cargo add`s
  `launchdarkly-sdk-transport` with the `hyper` feature so the
  programmatic proxy fragment's `HyperTransport` import resolves.

## Known non-binds

Ten shell fragments (`apex-server-sdk`, `go-server-sdk`,
`python-server-sdk`, `ruby-server-sdk` mac/windows pairs;
`rust-server-sdk` plain/authenticated pair) just `export VAR=…` or
`set VAR=…` rather than running a package manager. The
`shell-install` validator is a package-manager sniff
(npm/pnpm/yarn/pip/go/bower) followed by an artifact assertion
(node_modules, go.mod, etc.); a bare `export HTTPS_PROXY=…` doesn't
fit that contract — there's nothing to assert post-run, and routing
to a generic "run this in /bin/sh and check exit 0" runner would
silently green-light malformed bodies. These snippets are
byte-equality-verified through marker hashes, which is the strongest
check available for a fragment without a runtime success-line. Same
blocker as the `proxy-env-*` fragments documented in
`_sdk-docs-port-notes.md`.
