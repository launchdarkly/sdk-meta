---
id: go-server-sdk/sdk-docs/features/aimetrics/track-request
sdk: go-server-sdk
kind: reference
lang: go
description: Wrap an AI provider call with TrackRequest for the Go AI SDK.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only

---

```go
if cfg.Enabled() {

  response, err := tracker.TrackRequest(func(config *Config) (ProviderResponse, error) {

    // Make request to a provider, which automatically tracks metrics in LaunchDarkly.
    // When sending the request to a provider, use details from config.
    // For example, you can pass a model parameter (config.ModelParam) or messages (config.Messages).
    // Optionally, return response metadata for additional logging.

    return ProviderResponse{
      Usage: TokenUsage{
        Total: 1, // Token usage data
      },
      Metrics: Metrics{
        Latency: 10 * time.Millisecond, // Metrics data
      },
    }, nil
  })

} else {

  // Application path to take when the configuration is disabled

}
```
