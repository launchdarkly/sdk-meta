---
id: ios-client-sdk/sdk-docs/use-cocoapods-podfile
sdk: ios-client-sdk
kind: reference
lang: ruby
description: "Podfile in section \"Use CocoaPods\""
# TODO(snippet-bug): body's `LaunchDarklyObservability ~> 1.0`
# resolves to no released pod — observability has only 0.x tags on
# the swift-launchdarkly-observability repo, and the pod hasn't
# been published to cocoapods trunk at any 1.x version. The doc
# claim "requires iOS SDK v9.14+" is forward-looking. Fix in the
# snippet-bugs PR: drop the observability line, or update to a
# pod version that actually exists once one is published.
---

```ruby
use_frameworks!
target 'YourTargetName' do
  pod 'LaunchDarkly', '~> 9.0'
  # optional observability plugin, requires iOS SDK v9.14+
  pod 'LaunchDarklyObservability', '~> 1.0'
end
```
