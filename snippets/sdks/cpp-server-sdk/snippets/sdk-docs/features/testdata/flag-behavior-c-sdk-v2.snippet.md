---
id: cpp-server-sdk/sdk-docs/features/testdata/flag-behavior-c-sdk-v2
sdk: cpp-server-sdk
kind: reference
lang: c
description: Configuring test data flag behavior for the C server SDK v2.x (native).
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only-v2-c
---

```c
// This flag is true for the user key "example-user-key"
// and false for everyone else.
struct LDFlagBuilder *flag2 = LDTestDataFlag(td, "flag-key-456def");
LDFlagBuilderVariationForUserBoolean(flag2, "example-user-key", LDBooleanTrue);
LDFlagBuilderFallthroughVariationBoolean(flag2, LDBooleanFalse);
LDTestDataUpdate(td, flag2);


// This flag returns the string variation "green" for
// users who have the custom attribute "admin" with a
// value of true, and "red" for everyone else.

struct LDJSON *variations = LDNewArray();
LDArrayPush(variations, LDNewText("red"));
LDArrayPush(variations, LDNewText("green"));

struct LDFlagBuilder *flag3 = LDTestDataFlag(td, "flag-key-789ghi");
LDFlagBuilderVariations(flag3, variations);
LDFlagBuilderFallthroughVariation(flag3, 0);

struct LDFlagRuleBuilder *rule = LDFlagBuilderIfMatch(flag3, "admin", LDNewBool(LDBooleanTrue));
LDFlagRuleBuilderThenReturn(rule, 1);

LDTestDataUpdate(td, flag3);
```
