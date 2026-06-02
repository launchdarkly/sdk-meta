---
id: dotnet-server-sdk/sdk-docs/initialize-the-client-initialize-net-sdk-v8-10-with-asp-net-core-p2
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: "Initialize, .NET SDK v8.10+ with ASP.Net Core in section \"Initialize the client\""
# TODO(validate): ASP.NET Framework Global.asax fragment that nests
# `protected void Application_Start()` / `Application_End()` method
# definitions inside the wrappee body. Nested method definitions with
# `protected` modifiers aren't valid inside another method, so the
# csharp-syntax-only scaffold's wrapper makes this uncompilable.
# See _sdk-docs-port-notes.md.
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only
---

```csharp
  // In Global.asax.cs
  protected void Application_Start()
  {
      // Other application specific code.
      var client = new LdClient(Configuration.Builder("YOUR_SDK_KEY")
          .StartWaitTime(TimeSpan.FromSeconds(5))
          .Plugins(new PluginConfigurationBuilder().Add(ObservabilityPlugin.Builder()
              .WithServiceName("your-service-name")
              .WithServiceVersion("example-sha")
              .Build()))
          .Build());
  }
  
  protected void Application_End() {
      Observe.Shutdown();
  }

  // <!-- In Web.config -->
  // <system.webServer>
  //   <!-- Any existing content should remain and the following should be added. -->
  //   <modules>
  //   <add name="TelemetryHttpModule" type="OpenTelemetry.Instrumentation.AspNet.TelemetryHttpModule,
  //     OpenTelemetry.Instrumentation.AspNet.TelemetryHttpModule" preCondition="integratedMode,managedHandler" />
  //   </modules>
  // </system.webServer>
```
