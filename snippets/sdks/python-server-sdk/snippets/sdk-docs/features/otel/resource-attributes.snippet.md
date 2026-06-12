---
id: python-server-sdk/sdk-docs/features/otel/resource-attributes
sdk: python-server-sdk
kind: reference
lang: python
description: Programmatic OpenTelemetry resource attribute configuration (Python).
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only
---

```python
from opentelemetry.sdk.resources import Resource
from opentelemetry.sdk.trace import TracerProvider

resource = Resource.create({
    "launchdarkly.project_id": "YOUR_SDK_KEY"
})

provider = TracerProvider(resource=resource)
```
