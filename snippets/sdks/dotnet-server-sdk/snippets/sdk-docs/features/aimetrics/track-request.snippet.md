---
id: dotnet-server-sdk/sdk-docs/features/aimetrics/track-request
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: Wrap an AI provider call with TrackRequest for the .NET AI SDK.
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only

---

```csharp
if (tracker.Config.Enabled == true) {

  // Use Task.Run here only as an example. You can wrap any async provider call.
  var response = tracker.TrackRequest(Task.Run(() =>
    {
      // Make request to a provider, which automatically tracks metrics in LaunchDarkly.
      // When sending the request to a provider, use details from tracker.Config.
      // For example, you can pass tracker.Config.Model and tracker.Config.Messages.
      // Optionally, return response metadata for additional logging.
      //
      // CAUTION: If the call inside Task.Run throws an exception,
      // the SDK will re-throw that exception.

      return new Response
      {
        Usage = new Usage { Total = 1, Input = 1, Output = 1 }, // Token usage data
        Metrics = new Metrics { LatencyMs = 100 } // Metrics data
      };
    }
  ));

} else {

  // Application path to take when tracker.Config is disabled

}
```
