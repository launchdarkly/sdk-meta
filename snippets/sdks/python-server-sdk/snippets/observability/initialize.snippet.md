---
id: python-server-sdk/observability/initialize
sdk: python-server-sdk
kind: initialize
lang: python
file: python-server-sdk/observability/initialize.txt
description: Initialize python-server-sdk with observability plugin and emit a sample log/span.
---

```python
ldclient.set_config(Config(
    'SDK_KEY',
    # ... all existing options
    plugins=[
        ObservabilityPlugin(
            ObservabilityConfig(
                service_name="my-service-name",
                service_version="1.0.0",
            )
        )
    ]
))

# Record a log
observe.record_log("Custom log message", logging.INFO, {"custom": "value"})

# Start a span
with observe.start_span("operation-name", attributes={"custom": "value"}) as span:
    span.set_attribute("my-attribute", "my-value")
    # Your code here
```
