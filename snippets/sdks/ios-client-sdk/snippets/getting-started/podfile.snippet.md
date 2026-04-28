---
id: ios-client-sdk/getting-started/podfile
sdk: ios-client-sdk
kind: manifest
lang: ruby
file: Podfile
description: CocoaPods Podfile pulling the LaunchDarkly iOS SDK.
inputs:
  version:
    type: string
    description: SDK version. Defaults to '6.1.0' in gonfalon as a fallback when the async fetch hasn't completed.
    runtime-default: "6.1.0"
ld-application:
  slot: podfile
---

Install the LaunchDarkly SDK using [CocoaPods](https://cocoapods.org/) by creating a `Podfile`:

```ruby
target 'hello-swift' do
  pod 'LaunchDarkly', '{{ version }}'
end

```
