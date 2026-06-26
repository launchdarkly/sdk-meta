---
id: sdk-docs/features/otel/resource-attributes-collector
kind: reference
lang: yaml
description: OpenTelemetry collector resource processor setting the LaunchDarkly SDK key attribute.
# TODO(validate): no yaml validator runtime exists yet, and _shared snippets carry no sdk: field for a CI row to select. See _feature-docs-otel-port-notes.md.
---

```yaml
# excerpt from otel-collector-config.yaml
processors:
  resource:
    attributes:
      - key: launchdarkly.project_id
        value: YOUR_SDK_KEY
        action: upsert
```
