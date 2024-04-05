.PHONY: crawl products

# Crawl all the repos and update metadata.sqlite3 with the results.
# GITHUB_TOKEN must be set in the environment.
crawl:
	./scripts/crawl.sh metadata.sqlite3 metadata

# Generate all the JSON products.
products:
	./scripts/generate-products.sh

all: crawl products
