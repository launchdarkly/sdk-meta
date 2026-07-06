# Edge section port notes

Source: `ld-docs-private/fern/topics/sdk/edge/` — `akamai.mdx`,
`cloudflare/index.mdx`, `fastly.mdx`, `vercel.mdx` (the `index.mdx` overview
and `cloudflare/migration-1-to-2.mdx` are not ported; the migration page is
version-specific and a poor validation target).

These are the per-edge-SDK getting-started/reference pages (install ·
import · initialize · evaluate · example worker), distinct from the
`sdk-docs/features/*` edge snippets already landed by the features fan-out.
Snippets live at the top level of each edge SDK's `sdk-docs/` group.

## Validation: real type-check, not parse-only

Edge SDKs run in edge runtimes (Akamai EdgeWorkers, Cloudflare Workers,
Fastly Compute, Vercel Edge) that can't execute against a live LD
environment in CI, so the strongest available validation is type-checking
against the real published types. The new **`edge-tsc`** validator installs
the real edge SDK packages plus each runtime's ambient type packages and
runs `tsc --noEmit` with full module resolution and type-checking. The
existing parse-only `edge-ts` validator is left in place for the
feature-page fragments that use it.

Per-SDK scaffolds compose fragments so the real API is exercised:
- `*-toplevel` — resolves the import fragment against the real package.
- `*-init-runner` — supplies the documented import(s), then splices the
  init fragment so `init(...)` is checked against the real signature.
- `*-eval` — supplies an initialized client (prepended, or ambiently typed
  via a `_globals.d.ts` companion when the fragment carries its own
  imports), so `variation(...)` is checked against the real client API.
- cloudflare adds an `*-worker` scaffold for the exported-fetch-handler
  example.

Notes:
- The Cloudflare KV binding is typed as `Parameters<typeof init>[1]` rather
  than a standalone `@cloudflare/workers-types` `KVNamespace`, to avoid the
  structural mismatch between the workers-types version the SDK bundles and
  any version installed separately.
- The edge-tsc image installs with `--legacy-peer-deps`: the four edge SDKs
  pin overlapping ranges of the shared `js-server-sdk-common` peer; we only
  need each package's `.d.ts` on disk for type-checking.
- Not validated (rendered from canonical, but no runtime): Cloudflare's JSR
  install (`npx jsr add …`, not one of the shell-install package managers)
  and the `wrangler.toml` config block. The CommonJS `require(...)` import
  variants type-check trivially (require returns `any`); the TypeScript
  import variants get the real resolution check.

## CI

No matrix edit: the akamai/cloudflare/fastly/vercel rows already validate
`group: sdk-docs` with `key-type: none`, so the new snippets and the
`edge-tsc` validator are picked up automatically. Independent of the
OpenFeature PR (#531) — no shared harness files overlap.
