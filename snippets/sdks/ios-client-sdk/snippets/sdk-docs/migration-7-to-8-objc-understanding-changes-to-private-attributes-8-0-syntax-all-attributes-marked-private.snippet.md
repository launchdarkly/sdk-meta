---
id: ios-client-sdk/sdk-docs/migration-7-to-8-objc-understanding-changes-to-private-attributes-8-0-syntax-all-attributes-marked-private
sdk: ios-client-sdk
kind: reference
lang: objective-c
description: "8.0 syntax, all attributes marked private in section \"Understanding changes to private attributes\""
---

```objectivec
LDConfig *config = [[LDConfig alloc] initWithMobileKey:@"example-mobile-key"];
[config setAllContextAttributesPrivate:YES];
[LDClient startWithConfiguration:config context:context completion:nil];
```
