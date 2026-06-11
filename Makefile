.PHONY: help
help: #! Show this help message
	@echo 'Usage: make [target] ... '
	@echo ''
	@echo 'Targets:'
	@grep -h -F '#!' $(MAKEFILE_LIST) | grep -v grep | sed 's/:.*#!/:/' | column -t -s":"

# GITHUB_TOKEN must be set in the environment.
.PHONY: crawl
crawl: #! Crawl all the repos and update metadata.sqlite3 with the results
	./scripts/crawl.sh metadata.sqlite3 metadata

.PHONY: products
products: #! Generate all the JSON products
	./scripts/generate-products.sh

.PHONY: html
html: #! Generate SDK features HTML comparison pages (both views)
	mkdir -p _site
	cd tool && go run ./cmd/genhtml --view=by-sdk --output ../_site/by-sdk.html --data ../products
	cd tool && go run ./cmd/genhtml --view=by-feature --output ../_site/by-feature.html --data ../products
	cp _site/by-sdk.html _site/index.html

# Spec-support pipeline (genspecs). See scripts/generate-spec-support.sh.
#
# These use the same `cd tool && go run ./cmd/X --... ../products/...` pattern
# as the existing `html` target. `go run` inherits cwd from the shell, so we
# pass repo-root-relative paths with a `..` prefix.

.PHONY: spec-sync-repos
spec-sync-repos: #! Clone any missing SDK repo (plus sdk-specs and sdk-test-harness) and fast-forward existing ones
	cd tool && go run ./cmd/genspecs sync-repos --repos-json ../products/repos.json

.PHONY: spec-catalog
spec-catalog: #! Generate products/specs.json from the local sdk-specs checkout
	cd tool && go run ./cmd/genspecs catalog --out ../products/specs.json

.PHONY: spec-harness
spec-harness: #! Generate products/harness_signals.json from sdk-test-harness and per-SDK suppressions
	cd tool && go run ./cmd/genspecs harness \
	  --repos-json ../products/repos.json \
	  --out ../products/harness_signals.json

.PHONY: spec-judge
spec-judge: #! Run the LLM judge to populate products/spec_support.json (defaults to noop if ANTHROPIC_API_KEY is unset)
	cd tool && go run ./cmd/genspecs judge \
	  --specs-json ../products/specs.json \
	  --harness-json ../products/harness_signals.json \
	  --features-json ../products/features.json \
	  --feature-info-json ../products/feature_info.json \
	  --types-json ../products/types.json \
	  --names-json ../products/names.json \
	  --languages-json ../products/languages.json \
	  --repos-json ../products/repos.json \
	  --cache-dir ../tool/specs/.judge-cache \
	  --out ../products/spec_support.json

.PHONY: spec-html
spec-html: #! Render _site/spec-support*.html from the spec_support data
	mkdir -p _site
	cd tool && go run ./cmd/genspecs html \
	  --specs-json ../products/specs.json \
	  --support-json ../products/spec_support.json \
	  --types-json ../products/types.json \
	  --names-json ../products/names.json \
	  --out-dir ../_site

.PHONY: spec-support
spec-support: #! Run the full spec-support pipeline (sync-repos -> catalog -> harness -> judge -> html)
	./scripts/generate-spec-support.sh

.PHONY: all
all: crawl products html
