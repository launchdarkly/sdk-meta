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
    description: SDK version. Defaults to '11.1.2' as a fallback when gonfalon's async fetch from CocoaPods hasn't completed.
    runtime-default: "11.1.2"
ld-application:
  slot: podfile
---

Install the LaunchDarkly SDK using [CocoaPods](https://cocoapods.org/) by creating a `Podfile`:

```ruby
target 'hello-swift' do
  pod 'LaunchDarkly', '{{ version }}'
end

```
