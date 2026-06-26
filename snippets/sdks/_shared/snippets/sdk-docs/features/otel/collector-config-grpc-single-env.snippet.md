---
id: sdk-docs/features/otel/collector-config-grpc-single-env
kind: reference
lang: yaml
description: OpenTelemetry collector configuration sending traces to LaunchDarkly over gRPC (one environment).
# TODO(validate): no yaml validator runtime exists yet, and _shared snippets carry no sdk: field for a CI row to select. See _feature-docs-otel-port-notes.md.
---

```yaml
# The receivers specify how the Collector receives data.
# In this example, it receives data using the OpenTelemetry Protocol (OTLP) over gRPC and HTTP.
receivers:
  otlp:
    protocols:
      grpc:
      http:

# The exporters specify how the Collector sends data.
# In this example, it sends data to LaunchDarkly using gRPC.
exporters:
  otlp/jaeger:
    endpoint: jaeger-all-in-one:4317
    tls:
      insecure: true

  otlphttp:
    endpoint: https://otel.observability.app.launchdarkly.com:4318
  otlp:
    endpoint: https://otel.observability.app.launchdarkly.com:4317

# The processors specify how the Collector processes the trace data.
# In this example, it drops spans that aren't needed for guarded rollouts.
processors:
  # This processor sets the resource attribute to your SDK key.
  # If you are using a LaunchDarkly SDK, you do NOT need to include this.
  # If you are using an OpenTelemetry SDK, you do need to set the
  # resource attributes. You can set them here, set them in your application,
  # or set them using an environment variable.
  resource:
    attributes:
        - key: launchdarkly.project_id
          value: YOUR_SDK_KEY
          action: upsert

  # This filter drops all span events except for exceptions and feature flag evaluations
  # that were part of an active guarded rollout.
  # It is optional, but useful to limit the amount of data you send to LaunchDarkly.
  filter/launchdarkly-spanevents:
    error_mode: ignore
    traces:
      spanevent:
        - 'not ((name == "feature_flag" and attributes["feature_flag.result.reason.inExperiment"] == true) or name == "exception")'

  # Remove all spans that do not have an HTTP route or any span events remaining
  # after the previous filter has been applied
  filter/launchdarkly-spans:
    error_mode: ignore
    traces:
      span:
        - 'not (attributes["http.route"] != nil or Len(events) > 0)'

  batch:

extensions:
  health_check:

service:
  pipelines:
    # Add a new pipeline to send data to LaunchDarkly
    traces/ld:
      receivers: [otlp]
      processors:
        [
          filter/launchdarkly-spanevents,
          filter/launchdarkly-spans,
          batch,
        ]
      exporters: [otlphttp/launchdarkly]
    metrics:
      receivers: [otlp]
      processors: [resource]
      exporters: [otlphttp]
    logs:
      receivers: [otlp]
      processors: [resource]
      exporters: [otlphttp]
```
