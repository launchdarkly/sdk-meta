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

.PHONY: all
all: crawl products html
