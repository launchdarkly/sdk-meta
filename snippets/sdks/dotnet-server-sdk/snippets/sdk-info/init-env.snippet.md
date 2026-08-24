---
id: dotnet-server-sdk/sdk-info/init-env
sdk: dotnet-server-sdk
kind: init
lang: csharp
file: dotnet-server-sdk/init-env.txt
description: Client initialization snippet for dotnet-server-sdk reading the SDK key from an environment variable.
validation:
  scaffold: dotnet-server-sdk/scaffolds/init-runner
---

```csharp
using LaunchDarkly.Sdk;
using LaunchDarkly.Sdk.Server;

var builder = WebApplication.CreateBuilder(args);

// Reads your LaunchDarkly SDK key from the LAUNCHDARKLY_SDK_KEY
// environment variable, so one build can run in every environment.
var ldConfig = Configuration.Default(Environment.GetEnvironmentVariable("LAUNCHDARKLY_SDK_KEY"));
var client = new LdClient(ldConfig);

if (client.Initialized)
{
    // For onboarding purposes only we flush events as soon as
    // possible so we quickly detect your connection.
    // You don't have to do this in practice because events are automatically flushed.
    client.Flush();
    Console.WriteLine("*** SDK successfully initialized!\n");
}
else
{
    Console.WriteLine("*** SDK failed to initialize\n");
    Environment.Exit(1);
}
```
