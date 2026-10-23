---
id: sdk-docs/features/otel/resource-attributes-env-var
kind: reference
lang: bash
description: OTEL_RESOURCE_ATTRIBUTES environment variable carrying the LaunchDarkly SDK key.
# TODO(validate): bare shell export command; not a package-manager install, no harness fits, and _shared snippets carry no sdk: field for a CI row to select.
# Do not use a run-and-check-exit-0 runner: it would pass bodies that are not correct.
---

```bash
export OTEL_RESOURCE_ATTRIBUTES="launchdarkly.project_id=$YOUR_SDK_KEY"
```
