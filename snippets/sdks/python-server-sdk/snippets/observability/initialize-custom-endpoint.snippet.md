---
id: python-server-sdk/observability/initialize-custom-endpoint
sdk: python-server-sdk
kind: initialize
lang: python
file: python-server-sdk/observability/initialize-custom-endpoint.txt
description: Initialize python-server-sdk with the observability plugin against a custom observability domain (OBSERVABILITY_BASE_DOMAIN is substituted by the consumer) and emit a sample log/span.
validation:
  scaffold: python-server-sdk/scaffolds/init-runner-observability
  placeholders:
    SDK_KEY: LAUNCHDARKLY_SDK_KEY
---

```python
ldclient.set_config(Config(
    'SDK_KEY',
    # ... all existing options
    plugins=[
        ObservabilityPlugin(
            ObservabilityConfig(
                backend_url="https://pub.OBSERVABILITY_BASE_DOMAIN",
                otlp_endpoint="https://otel.OBSERVABILITY_BASE_DOMAIN:4317",
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
