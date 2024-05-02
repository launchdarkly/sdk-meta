# sdk-meta

This repo contains metadata related to LaunchDarkly SDKs. 

The data is intended for consumption by downstream products and services.

| Data Product                             | Description                                                     | Format   |
|------------------------------------------|-----------------------------------------------------------------|----------|
| [Database](./metadata.sqlite3)           | Database containing data from which other products are derived. | sqlite3  |
| [SDK Names](products/names.json)         | SDK friendly names for display.                                 | JSON     |
| [SDK Releases](products/releases.json)   | SDK major/minor releases with EOL dates.                        | JSON     |
| [SDK Types](products/types.json)         | SDK types for categorization.                                   | JSON     |
| [SDK Features](products/features.json)   | SDK features, including version introduced/deprecated.          | JSON     |
| [SDK Languages](products/languages.json) | Programming languages associated with SDKs.                     | JSON     |
| [SDK Repos](products/repos.json)         | SDK source repositories                                         | JSON     |


## structure

This repo contains an sqlite database containing a snapshot of SDK metadata
fetched from individual repos.

It also contains JSON files that are derived from the database. These are intended for
consumption by downstream products and services.

The JSON data products live in [`products`](./products) and the schemas for them live in [`schemas`](./schemas). 

## tooling

Data products can be validated against the schemas using `./scripts/ci/check-json-schemas.sh` on Linux.

Ensure that the JSON files are valid and formatted using `./scripts/ci/format-json.sh`.

To quickly create basic metadata for a new repo, run `./scripts/add-repo.sh <launchdarkly/repo-name>`. This will
clone the repo, check-in an `.sdk_metadata.json` file, and create a PR.

## versioning policy

The JSON products have an implicit 'v1' version.

The sqlite database has no versioning policy and may change schemas at any time.

If any JSON schema needs a breaking change, then we should 
probably handle it like so:
1. Add a new `products/v2` directory and `schemas/v2` directory as needed
2. Update `./scripts/generate-products.sh` to generate the new data products, and also keep generating the
old ones so they remain up-to-date.
3. Have new applications consume `/products/v2/whatever.json`
4. Eventually, when/if all consumers have migrated, stop generating the old products. Don't delete them though, 
just in case something is intermittently using them that we've forgotten about.

### what necessitates a new major version
- Changing the meaning of an existing field
- Changing the contents of an existing field in an incompatible way (like changing its type)
- Removing an existing field
- Probably other things that we haven't thought of

### what doesn't necessitate a new major version
- Adding a new field


## consumers

Consumers of SDK metadata include:
- [LaunchDarkly Docs](https://docs.launchdarkly.com/)

Are you a consumer? Add a link here to receive communications when new features/changes are coming.
