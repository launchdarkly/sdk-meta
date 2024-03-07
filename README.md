# sdk-meta

This repo contains metadata related to LaunchDarkly SDKs. 

The data is intended for consumption by downstream products and services.

| Data Product                           | Description                                            |
|----------------------------------------|--------------------------------------------------------|
| [SDK List](./data/sdks.json)           | Authoritative list of LaunchDarkly SDK IDs.            |
| [SDK Names](./data/names.json)         | SDK friendly names for display.                        |
| [SDK Types](./data/types.json)         | SDK types for categorization.                          |
| [SDK Features](./data/features.json)   | SDK features, including version introduced/deprecated. |
| [SDK Languages](./data/languages.json) | Programming languages associated with SDKs.            |
| [SDK Repos](./data/repos.json)         | SDK source repositories                                |


## structure

This repo is essentially a JSON database hosted on Github. 

The "tables" live in [`data`](./data) and the schemas for those tables live in [`schemas`](./schemas). When adding
a new table, ensure it has a corresponding schema.

Data can be validated against the schemas using `./scripts/ci/check-json-schemas.sh` on Linux.

Ensure that the JSON files are valid and formatted using `./scripts/ci/format-json.sh`.

Finally, it's important that each SDK listed in `sdks.json` has a corresponding entry in each table, as this is
what downstream tools will expect. 

This is enforced with `./scripts/ci/sdk-consistency.sh`. 
