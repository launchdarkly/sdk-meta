# Agent guide — launchdarkly/sdk-meta

This repo is a monorepo. For most work the first thing to read is the per-subtree guide, not this file.

## Where to start

- **Snippets work** (authoring `.snippet.md` files, adding scaffolds, extending validators, editing the CI matrix): read `snippets/CLAUDE.md`. It's auto-loaded when your cwd is under `snippets/`. Don't author snippet content in consumer repos (gonfalon, ld-docs) — those only carry markers that point back here.
- **Non-snippets work**: read `CLAUDE.md` at the repo root if it exists. Otherwise read the README under the relevant subtree.

## Top-level surfaces

| Path | Purpose |
|---|---|
| `snippets/` | Canonical home of every SDK snippet body, plus the CLI (`cmd/snippets`), per-language validator harnesses (`validators/`), and the validation CI workflow. **See `snippets/CLAUDE.md`.** |
| `api/` | OpenAPI specs and generated artifacts (Go side). |
| `api-js/` | JS/TS package built from the API specs. Released by release-please alongside `snippets`. |
| `metadata/` | SDK metadata source (versions, support matrix, end-of-life dates). |
| `products/` | Generated product/SDK-info JSON consumed by docs and the LD dashboard. |
| `backfill/` | One-off scripts for back-filling historical metadata. |
| `schemas/` | JSON schemas for the metadata + products outputs. |
| `scripts/` | Repo-level shell tooling. |
| `tool/` | Go tooling for metadata generation. |
| `.github/workflows/` | CI — `snippets-validate.yml` (snippets matrix), `release-please.yml` (monorepo releases). |

## Conventions across the repo

- Conventional Commits with a scope: `feat(snippets): …`, `fix(metadata): …`, `chore: …`. release-please drives changelogs and version bumps from these.
- Two release-please packages: `api-js` and `snippets`. Each ships its own tag (`api-js/vX.Y.Z`, `snippets/vX.Y.Z`) and changelog.
- For snippet PRs: branch `rlamb/sdk-NNNN/desc`, ticket in branch + PR body (not PR title), sequential PRs, base on `main` once predecessor has merged.
