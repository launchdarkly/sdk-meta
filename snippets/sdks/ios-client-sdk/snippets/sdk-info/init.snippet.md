---
id: ios-client-sdk/sdk-info/init
sdk: ios-client-sdk
kind: init
lang: swift
file: ios-client-sdk/init.txt
description: Client initialization snippet for ios-client-sdk.
validation:
  scaffold: ios-client-sdk/scaffolds/init-runner
  placeholders:
    YOUR_MOBILE_KEY: LAUNCHDARKLY_MOBILE_KEY
---

```swift
import LaunchDarkly

// This is your mobile key.
let config = LDConfig(mobileKey: "YOUR_MOBILE_KEY", autoEnvAttributes: .enabled)

// A "context" is a data object representing users, devices, organizations, and other entities.
let contextBuilder = LDContextBuilder(key: "EXAMPLE_CONTEXT_KEY")
guard case .success(let context) = contextBuilder.build()
else { return }

LDClient.start(config: config, context: context, startWaitSeconds: 5) { timedOut in
    if timedOut {
        print("SDK didn't initialize in 5 seconds.  SDK is still running and trying to get latest flags.")
    } else {
        print("SDK successfully initialized with the latest flags")
    }
}

print("SDK started.")
```
