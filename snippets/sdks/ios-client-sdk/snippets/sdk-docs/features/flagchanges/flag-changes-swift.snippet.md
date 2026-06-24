---
id: ios-client-sdk/sdk-docs/features/flagchanges/flag-changes-swift
sdk: ios-client-sdk
kind: reference
lang: swift
description: Flag change observation example for iOS (Swift).
validation:
  scaffold: ios-client-sdk/scaffolds/swift-syntax-only

---

```swift
let flagKey = "example-flag-key"
let flagObserverOwner = flagKey as LDObserverOwner

let client = LDClient.get()!

// Observe a specific flag
client.observe(keys: [flagKey], owner: flagObserverOwner, handler: { [weak self] changedFlags in
    if let changedFlag = changedFlags[flagKey] {
        let newValue = client.boolVariation(forKey: flagKey, defaultValue: false)
        print("Flag \(flagKey) changed to: \(newValue)")
    }
})

client.stopObserving(owner: flagObserverOwner)

// Observe when flags are unchanged
client.observeFlagsUnchanged(owner: self) {
    client.stopObserving(owner: self as LDObserverOwner)
}

// Observe all flag changes
client.observeAll(owner: self) { [weak self] changedFlags in
    for (key, changedFlag) in changedFlags {
        print("\(key) changed from \(changedFlag.oldValue) to \(changedFlag.newValue)")
    }
    client.stopObserving(owner: self as LDObserverOwner)
}
```
