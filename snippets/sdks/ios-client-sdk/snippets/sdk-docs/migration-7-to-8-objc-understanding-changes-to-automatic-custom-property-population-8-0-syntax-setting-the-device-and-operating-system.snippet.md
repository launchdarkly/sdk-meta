---
id: ios-client-sdk/sdk-docs/migration-7-to-8-objc-understanding-changes-to-automatic-custom-property-population-8-0-syntax-setting-the-device-and-operating-system
sdk: ios-client-sdk
kind: reference
lang: objective-c
description: "8.0 syntax, setting the device and operating system in section \"Understanding changes to automatic custom property population\""
---

```objectivec
LDContextBuilder *builder = [[LDContextBuilder alloc] initWithKey:@"example-context-key"];
[builder kindWithKind:@"device"];
[builder trySetValueWithName:@"device" value:[LDValue ofString:@"device"]];
[builder trySetValueWithName:@"os" value:[LDValue ofString:@"os"]];
LDContext *context = builder.build.success;
```
