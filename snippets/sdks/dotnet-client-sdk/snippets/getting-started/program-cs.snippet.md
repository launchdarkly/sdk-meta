---
id: dotnet-client-sdk/getting-started/program-cs
sdk: dotnet-client-sdk
kind: hello-world
lang: csharp
file: Program.cs
description: Hello-world program that initializes the .NET client SDK and evaluates a feature flag.
inputs:
  mobileKey:
    type: mobile-key
    description: Mobile key baked into the rendered source. Validation reads LAUNCHDARKLY_MOBILE_KEY at runtime.
  featureKey:
    type: flag-key
    description: Default flag key baked into the rendered source. Validation reads LAUNCHDARKLY_FLAG_KEY at runtime.
ld-application:
  slot: program-cs
validation:
  runtime: dotnet-client
  requirements: Launchdarkly.ClientSdk
---

Open the file `Program.cs` and add the following code:

```csharp
using LaunchDarkly.Sdk;
using LaunchDarkly.Sdk.Client;

var context = Context.New("context-key-123abc");
var timeSpan = TimeSpan.FromSeconds(10);
var client = LdClient.Init(
  Configuration.Default("{{ mobileKey }}", ConfigurationBuilder.AutoEnvAttributes.Enabled),
  context,
  timeSpan
);

if (client.Initialized)
{
    Console.WriteLine("SDK successfully initialized!");
}
else
{
    Console.WriteLine("SDK failed to initialize");
    Environment.Exit(1);
}

var flagValue = client.BoolVariation("{{ featureKey }}", false);

Console.WriteLine(string.Format("The '{{ featureKey }}' feature flag evaluates to {0}.", flagValue));

// Here we ensure that the SDK shuts down cleanly and has a chance to deliver analytics
// events to LaunchDarkly before the program exits. If analytics events are not delivered,
// the context properties and flag usage statistics will not appear on your dashboard. In
// a normal long-running application, the SDK would continue running and events would be
// delivered automatically in the background.
client.Dispose();
```
