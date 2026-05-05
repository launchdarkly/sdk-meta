---
id: dotnet-server-sdk/sdk-docs/initialize-the-client-initialize-net-sdk-v8-10-with-asp-net-core-p1
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: "Initialize, .NET SDK v8.10+ with ASP.Net Core in section \"Initialize the client\""
# Bucket C: ASP.NET Core init fragment. Requires Microsoft.AspNetCore +
# LaunchDarkly.Observability + ObservabilityPlugin types that the
# csharp-syntax-only scaffold doesn't pull in, plus the body references
# the bare identifier `args` which is only in scope inside the
# top-level Main(string[] args). See _sdk-docs-port-notes.md.
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
