---
id: sdk-docs/features/otel/resource-attributes-env-var
kind: reference
lang: bash
description: OTEL_RESOURCE_ATTRIBUTES environment variable carrying the LaunchDarkly SDK key.
# TODO(validate): bare shell export command; not a package-manager install, no harness fits, and _shared snippets carry no sdk: field for a CI row to select. See _feature-docs-otel-port-notes.md.
---

```bash
export OTEL_RESOURCE_ATTRIBUTES="launchdarkly.project_id=$YOUR_SDK_KEY"
```
