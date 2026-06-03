---
id: ios-client-sdk/sdk-docs/initialize-the-client-ios-sdk-v9-x-swift
sdk: ios-client-sdk
kind: reference
lang: swift
description: "iOS SDK v9.x (Swift) in section \"Initialize the client\""
# TODO(snippet-bug): body declares `let config = LDConfig(...)` then
# mutates `config.plugins = [...]`. LDConfig is a Swift struct
# (value type), so `let` makes it immutable and the assignment
# fails to compile. Fix in the snippet-bugs PR: change `let config`
# to `var config`.
---

```swift
  let config = LDConfig(mobileKey: "example-mobile-key", autoEnvAttributes: .enabled)
  // optional observability plugin, requires iOS SDK v9.14+
  config.plugins = [
        Observability(
            options: .init(
                resourceAttributes: [
                    "my-attribute": .string("new-value")
                ],
                isDebug: true,
                logs: .enabled,
                traces: .enabled,
                metrics: .enabled
            )
        )
    ]

  // You'll need this context later, but you can ignore it for now.
  let contextBuilder = LDContextBuilder(key: "example-context-key")
  guard case .success(let context) = contextBuilder.build()
  else { return }

  LDClient.start(config: config, context: context, startWaitSeconds: 5) { timedOut in
      if timedOut {
          // Client may not have the most recent flags for the configured context
      } else {
          // Client has received flags for the configured context
      }
  }
```
