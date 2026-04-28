---
id: dotnet-server-sdk/getting-started/program-cs
sdk: dotnet-server-sdk
kind: hello-world
lang: csharp
file: Program.cs
description: Hello-world program that initializes the .NET server SDK and watches a feature flag.
inputs:
  featureKey:
    type: flag-key
    description: Default flag key baked into the rendered source. Validation reads LAUNCHDARKLY_FLAG_KEY at runtime.
ld-application:
  slot: program-cs
validation:
  runtime: dotnet-server
  requirements: LaunchDarkly.ServerSdk
---

Open the file `Program.cs` and add the following code:

```csharp
using System;
  using System.Threading.Tasks;
  using LaunchDarkly.Sdk;
  using LaunchDarkly.Sdk.Server;

  namespace HelloDotNet
  {
      class Hello
      {
          public static void ShowBanner(){
              Console.WriteLine(
  @"            ██
            ██
        ████████
           ███████
  ██ LAUNCHDARKLY █
           ███████
        ████████
            ██
          ██
  ");
          }

          static void Main(string[] args)
          {
              bool CI = Environment.GetEnvironmentVariable("CI") != null;

              string SdkKey = Environment.GetEnvironmentVariable("LAUNCHDARKLY_SDK_KEY");

              // Set FeatureFlagKey to the feature flag key you want to evaluate.
              string FeatureFlagKey = "{{ featureKey }}";

              if (string.IsNullOrEmpty(SdkKey))
              {
                  Console.WriteLine("*** Please set LAUNCHDARKLY_SDK_KEY environment variable to your LaunchDarkly SDK key first\n");
                  Environment.Exit(1);
              }

              var ldConfig = Configuration.Default(SdkKey);

              var client = new LdClient(ldConfig);

              if (client.Initialized)
              {
                  Console.WriteLine("*** SDK successfully initialized!\n");
              }
              else
              {
                  Console.WriteLine("*** SDK failed to initialize\n");
                  Environment.Exit(1);
              }

              // Set up the evaluation context. This context should appear on your LaunchDarkly contexts
              // dashboard soon after you run the demo.
              var context = Context.Builder("example-user-key")
                  .Name("Sandy")
                  .Build();

              if (Environment.GetEnvironmentVariable("LAUNCHDARKLY_FLAG_KEY") != null)
              {
                  FeatureFlagKey = Environment.GetEnvironmentVariable("LAUNCHDARKLY_FLAG_KEY");
              }

              var flagValue = client.BoolVariation(FeatureFlagKey, context, false);

              Console.WriteLine(string.Format("*** The {0} feature flag evaluates to {1}.\n",
                  FeatureFlagKey, flagValue));

              if (flagValue)
              {
                  ShowBanner();
              }

              client.FlagTracker.FlagChanged += client.FlagTracker.FlagValueChangeHandler(
                  FeatureFlagKey,
                  context,
                  (sender, changeArgs) => {
                      Console.WriteLine(string.Format("*** The {0} feature flag evaluates to {1}.\n",
                      FeatureFlagKey, changeArgs.NewValue));

                      if (changeArgs.NewValue.AsBool) ShowBanner();
                  }
              );

              if(CI) Environment.Exit(0);

              Console.WriteLine("*** Waiting for changes \n");

              Task waitForever = new Task(() => {});
              waitForever.Wait();
          }
      }
  }
```
