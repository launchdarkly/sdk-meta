---
id: dotnet-server-sdk/sdk-docs/initialize-the-client-initialize-net-sdk-v8-10-with-asp-net-core-p1
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: "Initialize, .NET SDK v8.10+ with ASP.Net Core in section \"Initialize the client\""
# TODO(scaffold): needs a dedicated csharp-aspnet-core-syntax-only
# scaffold that (a) pulls Microsoft.AspNetCore.App + LaunchDarkly.Observability
# as requirements, (b) provides a top-level-statement Program.cs
# context where `args` is in scope (current csharp-syntax-only is a
# class-bodied console app), and (c) exposes PluginConfigurationBuilder
# + ObservabilityPlugin types. Once that scaffold lands, set
# `validation.scaffold:` here to point at it.
---

```csharp
  var builder = WebApplication.CreateBuilder(args);

  var config = Configuration.Builder("YOUR_SDK_KEY")
    .StartWaitTime(TimeSpan.FromSeconds(5))
    .Plugins(new PluginConfigurationBuilder()
        .Add(ObservabilityPlugin.Builder(builder.Services)
            .WithServiceName("your-service-name")
            .WithServiceVersion("example-sha")
            .Build()
        )
    ).Build();
  
  var client = new LdClient(config);

  // Client must be constructed before the web application.
  var app = builder.Build();
```
