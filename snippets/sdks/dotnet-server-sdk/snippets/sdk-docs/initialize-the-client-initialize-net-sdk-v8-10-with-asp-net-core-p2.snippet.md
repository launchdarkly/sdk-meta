---
id: dotnet-server-sdk/sdk-docs/initialize-the-client-initialize-net-sdk-v8-10-with-asp-net-core-p2
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: "Initialize, .NET SDK v8.10+ with ASP.Net Core in section \"Initialize the client\""
# TODO(scaffold): needs a dedicated csharp-aspnet-framework-syntax-only
# scaffold (legacy ASP.NET Framework, not Core) whose Program-equivalent
# is a `class : HttpApplication` so the body's `protected void
# Application_Start()` / `Application_End()` resolve as class members
# rather than nested method definitions. Requires
# System.Web + LaunchDarkly.Observability as references. Once that
# scaffold lands, set `validation.scaffold:` here to point at it.
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
