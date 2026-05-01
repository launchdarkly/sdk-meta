---
id: ios-client-sdk/sdk-docs/migration-5-to-6-objc-understanding-the-new-ldvalue-type-6-0-syntax
sdk: ios-client-sdk
kind: reference
lang: objective-c
description: "6.0 syntax in section \"Understanding the new LDValue type\""
---

```objectivec
LDValue *nullValue = [LDValue ofNull];
LDValue *boolValue = [LDValue ofBool:YES];
LDValue *numericValue = [LDValue ofNumber:@5.5];
LDValue *stringValue = [LDValue ofString:@"beta_testers"];
LDValue *complexValue = [LDValue ofDict:@{@"groups": [LDValue ofArray:@[[LDValue ofBool:YES]]]}];
```
