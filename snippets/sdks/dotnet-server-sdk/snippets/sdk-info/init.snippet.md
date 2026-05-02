---
id: dotnet-server-sdk/sdk-info/init
sdk: dotnet-server-sdk
kind: init
lang: csharp
file: dotnet-server-sdk/init.txt
description: Client initialization snippet for dotnet-server-sdk.
validation:
  scaffold: dotnet-server-sdk/scaffolds/init-runner
  placeholders:
    YOUR_SDK_KEY: LAUNCHDARKLY_SDK_KEY
---

```csharp
using LaunchDarkly.Sdk;
using LaunchDarkly.Sdk.Server;

var builder = WebApplication.CreateBuilder(args);

// This is your LaunchDarkly SDK key.
// Never hardcode your SDK key in production.
var ldConfig = Configuration.Default("YOUR_SDK_KEY");
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
