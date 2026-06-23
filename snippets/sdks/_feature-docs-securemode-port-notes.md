# Port notes: /sdk/features/secure-mode

Source: `ld-docs-private` `fern/topics/sdk/features/secure-mode.mdx`.
19 code blocks extracted into `sdk-docs/features/securemode/` snippets
across 15 SDKs. All 19 are bound to validators.

## Content corrections

None. Every body is verbatim from the MDX. All API claims were checked
against the local SDK repos before binding (`SecureModeHash` /
`secureModeHash` / `secure_mode_hash` signatures on .NET, Go
(LDClient + LDScopedClient), Haskell, Java, Node, PHP, Python, Ruby,
Rust v2/v3, and the js-core edge SDKs) and all matched the published
samples.

## Validation routing added in this port

- `haskell-server-sdk/scaffolds/haskell-syntax-only-expr` — the
  Haskell sample (`secureModeHash client context`) is a bare pure
  expression (`Client -> Context -> Text`). It is neither a monadic
  statement (so the `haskell-syntax-only` do-block rejects it) nor a
  top-level declaration (so `haskell-syntax-only-toplevel` would not
  parse it). The new scaffold splices the body as the right-hand side
  of a module-scope binding, with the same module-scope `client` /
  `context` stubs as the toplevel variant.
- `rust-server-sdk/scaffolds/rust-syntax-only-v2` — the Rust v2.x
  sample binds `secure_mode_hash` to a `String`, but the validator's
  Cargo project pulls the latest (3.x) crate where the method returns
  `Result<String, String>`. The scaffold compiles v2-era bodies
  against a stub client mirroring the 2.x surface, following the same
  approach as the beta `User` stub in `rust-syntax-only`.
- Rust validator image bumped from `rust:1.85` to `rust:1.94` — the
  3.x `launchdarkly-server-sdk` crate declares `rust-version = 1.93.0`
  as its MSRV, so the v3 sample (and any snippet compiled against the
  latest crate) needs the newer toolchain.

## Known non-binds

None. The page has no iOS Objective-C (or other scaffold-less) blocks.
