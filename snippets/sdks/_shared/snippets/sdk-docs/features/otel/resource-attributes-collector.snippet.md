---
id: sdk-docs/features/otel/resource-attributes-collector
kind: reference
lang: yaml
description: OpenTelemetry collector resource processor setting the LaunchDarkly SDK key attribute.
# TODO(validate): no yaml validator runtime exists yet, and _shared snippets carry no sdk: field for a CI row to select.
# To validate, add a yaml parse harness (or otelcol validate) and CLI support for snippets with no sdk: field.
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
