---
id: dotnet-client-sdk/sdk-info/init
sdk: dotnet-client-sdk
kind: init
lang: csharp
file: dotnet-client-sdk/init.txt
description: Client initialization snippet for dotnet-client-sdk.
---

```csharp
using LaunchDarkly.Sdk;
using LaunchDarkly.Sdk.Client;

// A "context" is a data object representing users, devices, organizations, and
// other entities. You'll need this to initialize the client.
Context context = Context.New("EXAMPLE_CONTEXT_KEY");

// This is your mobile key.
var timeSpan = TimeSpan.FromSeconds(5);
var client = await LdClient.InitAsync("YOUR_MOBILE_KEY", ConfigurationBuilder.AutoEnvAttributes.Enabled, context, timeSpan);

if (client.Initialized)
{
    Console.WriteLine("SDK successfully initialized!");
}
```
